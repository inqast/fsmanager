package authenticator

import (
	"github.com/inqast/fsmanager/internal/telegram/grpc"
)

type Authenticator struct {
	service *grpc.Service
}

func New(
	client *grpc.Service,
) *Authenticator {
	return &Authenticator{
		service: client,
	}
}
