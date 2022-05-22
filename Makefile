GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.17","$(shell printf "$(GO_VERSION_SHORT)\n1.17" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.17. Found: $(GO_VERSION_SHORT))
endif

export GO111MODULE=on

SERVICE_PATH=inqast/fsmanager

.PHONY: run
run:
	go run cmd/server/main.go

.PHONY: build
build:
	go mod download && CGO_ENABLED=0  go build \
		-tags='no_mysql no_sqlite3' \
		-ldflags=" \
			-X 'github.com/$(SERVICE_PATH)/internal/config.version=$(VERSION)' \
			-X 'github.com/$(SERVICE_PATH)/internal/config.commitHash=$(COMMIT_HASH)' \
		" \
		-o ./bin/service$(shell go env GOEXE) ./cmd/server/main.go

.PHONY: build-telegram
build-telegram:
	go mod download && CGO_ENABLED=0  go build \
		-tags='no_mysql no_sqlite3' \
		-ldflags=" \
			-X 'github.com/$(SERVICE_PATH)/internal/config.version=$(VERSION)' \
			-X 'github.com/$(SERVICE_PATH)/internal/config.commitHash=$(COMMIT_HASH)' \
		" \
		-o ./bin/telegram$(shell go env GOEXE) ./cmd/telegram/main.go

protobuf:
	protoc --proto_path=./api \
		--go_out=pkg/api \
		--go_opt=paths=source_relative \
		--go-grpc_out=pkg/api \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out pkg/api \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		--openapiv2_out pkg/api \
		./api/api.proto
