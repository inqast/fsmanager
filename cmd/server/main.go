package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/inqast/fsmanager/internal/config"

	"github.com/inqast/fsmanager/internal/app"
	"github.com/inqast/fsmanager/internal/db"
	"github.com/inqast/fsmanager/internal/repository"
	"github.com/inqast/fsmanager/pkg/api"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfigFromFile()

	db.Migrate(cfg)

	adp, err := db.New(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	newServer := app.New(repository.New(adp))
	lis, err := net.Listen("tcp", cfg.Grpc.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterFamilySubServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
}
