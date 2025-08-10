package config

import (
	"fmt"
	"time"
)

type ServerConfig struct {
	Host string
	Port string
}

type StoreConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

type JWTConfig struct {
	DurationMin time.Duration
	SecretKey   []byte
}

func LoadJWTConfig() *JWTConfig {
	return &JWTConfig{
		DurationMin: time.Duration(GetIntOrFatal(JWTTimeLimitMin)) * time.Minute,
		SecretKey:   []byte(GetValueOrFatal(JWTSecretKey)),
	}
}

// Получить конфиг сервера из .env файла
func LoadServerConfig() *ServerConfig {
	return &ServerConfig{
		Host: GetValueOrFatal(ServerHost),
		Port: GetValueOrFatal(ServerPort),
	}
}

func LoadStoreConfig() *StoreConfig {
	return &StoreConfig{
		Host:     GetValueOrFatal(DBHost),
		Port:     GetValueOrFatal(DBPort),
		User:     GetValueOrFatal(DBUser),
		Password: GetValueOrFatal(DBPassword),
	}
}

func (s *StoreConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=auth sslmode=disable",
		s.Host,
		s.Port,
		s.User,
		s.Password,
	)
}
