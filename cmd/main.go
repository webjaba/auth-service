package main

import (
	"auth-service/internal/app"
	"auth-service/internal/pkg/configuration"
	"log"
)

func main() {
	cfg := configuration.ServerConfig{
		Host: "localhost",
		Port: "50001",
	}

	service := app.New()

	if err := service.Start(cfg); err != nil {
		log.Fatalf("%s", err)
	}
}
