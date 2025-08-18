package store

import (
	"auth-service/internal/pkg/config"
	"log"
	"time"

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

	db.SetConnMaxIdleTime(time.Duration(cfg.MaxIdleTime) * time.Minute)
	db.SetMaxIdleConns(cfg.MaxIdleCons)
	db.SetMaxOpenConns(cfg.MaxConns)

	return &Store{db: db}
}

func (s *Store) Close() {
	s.db.Close()
}
