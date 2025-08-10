package store

import (
	"auth-service/internal/pkg/domain"
	"auth-service/internal/pkg/encrypt"
)

type IUser interface {
	InsertUser(user domain.User) (int, error)
}

func (s *Store) InsertUser(user domain.User) (int, error) {
	stmt, err := s.db.PrepareNamed(
		"INSERT INTO users (username, password) VALUES (:username, :password) RETURNING id",
	)
	if err != nil {
		return -1, err
	}

	user.Password, err = encrypt.EncryptPassword(user.Password)
	if err != nil {
		return -1, err
	}

	err = stmt.Get(&user.ID, user)
	if err != nil {
		return -1, err
	}
	return user.ID, nil
}
