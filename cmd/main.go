package main

import (
	"auth-service/cmd/configurations"
	"auth-service/internal/app"
	"auth-service/internal/pkg/store"

	"log"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()

	store := store.MustNew(configurations.StoreConfig)

	service := app.New(server, store)

	if err := service.Start(configurations.ServiceConfig); err != nil {
		log.Fatalf("%s", err)
	}
}
