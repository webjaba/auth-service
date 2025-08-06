package configuration

import "fmt"

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

func (s *StoreConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=auth sslmode=disable",
		s.Host,
		s.Port,
		s.User,
		s.Password,
	)
}
