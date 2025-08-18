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
	Host        string
	Port        string
	User        string
	Password    string
	MaxIdleCons int
	MaxIdleTime int
	MaxConns    int
}

type JWTConfig struct {
	Duration  time.Duration
	Leeway    time.Duration
	SecretKey []byte
}

func LoadJWTConfig() *JWTConfig {
	return &JWTConfig{
		Leeway:    time.Duration(GetIntOrFatal(JWTLeeway)) * time.Second,
		Duration:  time.Duration(GetIntOrFatal(JWTTimeLimitMin)) * time.Minute,
		SecretKey: []byte(GetValueOrFatal(JWTSecretKey)),
	}
}

// Получить конфиг сервера из .env файла
func LoadServerConfig() *ServerConfig {
	return &ServerConfig{
		Host: GetValueOrFatal(ServerHost),
		Port: GetValueOrFatal(ServerPort),
	}
}

// Получить конфиг стора из .env файла
func LoadStoreConfig() *StoreConfig {
	return &StoreConfig{
		Host:        GetValueOrFatal(DBHost),
		Port:        GetValueOrFatal(DBPort),
		User:        GetValueOrFatal(DBUser),
		Password:    GetValueOrFatal(DBPassword),
		MaxIdleCons: GetIntOrFatal(MaxIdleConns),
		MaxIdleTime: GetIntOrFatal(MaxIdleTime),
		MaxConns:    GetIntOrFatal(MaxConns),
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
