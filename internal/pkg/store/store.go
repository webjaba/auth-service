package store

import (
	"auth-service/internal/pkg/config"
	"log"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type IStore interface {
	IUser
}

type Store struct {
	db *sqlx.DB
}

func MustNew(cfg *config.StoreConfig) *Store {
	db, err := sqlx.Connect("postgres", cfg.DSN())
	if err != nil {
		log.Fatalf("db connection error: %s", err)
	}

	return &Store{db: db}
}
