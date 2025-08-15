package main

import (
	"auth-service/internal/app"
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/encrypt"
	"auth-service/internal/pkg/store"

	"log"

	"google.golang.org/grpc"
)

func main() {
	config.MustLoadEnv("")

	jwt := encrypt.NewJWT(config.LoadJWTConfig())

	server := grpc.NewServer()

	store := store.MustNew(config.LoadStoreConfig())

	service := app.New(server, store, jwt)

	if err := service.Start(config.LoadServerConfig()); err != nil {
		log.Fatalf("%s", err)
	}
}
