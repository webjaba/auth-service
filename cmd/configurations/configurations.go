package configurations

import "auth-service/internal/pkg/configuration"

var (
	ServiceConfig = &configuration.ServerConfig{
		Host: "localhost",
		Port: "50011",
	}

	StoreConfig = &configuration.StoreConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
	}
)
