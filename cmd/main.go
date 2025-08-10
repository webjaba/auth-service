package main

import (
	"auth-service/internal/app"
	"auth-service/internal/pkg/config"
	jwttoken "auth-service/internal/pkg/jwt_token"
	"auth-service/internal/pkg/store"

	"log"

	"google.golang.org/grpc"
)

func main() {
	config.MustLoadEnv("")

	jwtManager := jwttoken.NewManager(config.LoadJWTConfig())

	server := grpc.NewServer()

	store := store.MustNew(config.LoadStoreConfig())

	service := app.New(server, store, jwtManager)

	if err := service.Start(config.LoadServerConfig()); err != nil {
		log.Fatalf("%s", err)
	}
}
