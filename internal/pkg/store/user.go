package store

import (
	"auth-service/internal/pkg/domain"
)

type IUser interface {
	InsertUser(user *domain.User) error
}

func (s *Store) InsertUser(user *domain.User) error {
	stmt, err := s.db.PrepareNamed(
		"INSERT INTO users (username, password) VALUES (:username, :password) RETURNING id",
	)
	if err != nil {
		return err
	}

	err = stmt.Get(&user.ID, user)
	if err != nil {
		return err
	}
	return nil
}
