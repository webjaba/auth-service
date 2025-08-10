package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	errEmptyValue = errors.New("empty value")
)

const (
	JWTSecretKey    = "JWT_SECRET_KEY"
	JWTTimeLimitMin = "JWT_TIME_LIMIT_MIN"

	ServerHost = "SERVER_HOST"
	ServerPort = "SERVER_PORT"

	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
	DBUser     = "DB_USER"
	DBPassword = "DB_PASSWORD"
)

func MustLoadEnv(path string) {
	if path == "" {
		path = ".env"
	}
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("error load .env file: %s", err)
	}
}

func GetValueOrFatal(key string) string {
	v, err := getValue(key)
	if err != nil {
		log.Fatalf("error load '%s' value from .env: %s", key, err)
	}
	return v
}

func MustGetValue(key string) string {
	v, _ := getValue(key)
	return v
}

func getValue(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return "", errEmptyValue
	}
	return v, nil
}

func GetIntOrFatal(key string) int {
	i, err := strconv.Atoi(GetValueOrFatal(key))
	if err != nil {
		log.Fatalf("load '%s' from config error: %s", key, err)
	}
	return i
}
