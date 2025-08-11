package store

import (
	"auth-service/internal/pkg/domain"
	"auth-service/internal/pkg/encrypt"
	"errors"
	"strings"
)

var (
	errInvalidUsername = errors.New("invalid username")
)

type IUser interface {
	InsertUser(user domain.User) (int, error)
	GetUserByID(id int) (*domain.User, error)
}

func (s *Store) InsertUser(user domain.User) (int, error) {
	if ok := sqlInjectionCheck(user.Username); !ok {
		return 0, errInvalidUsername
	}
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

func (s *Store) GetUserByID(id int) (*domain.User, error) {
	user := &domain.User{}
	err := s.db.Get(user, "SELECT username, password FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func sqlInjectionCheck(value string) bool {
	incorrectSymbols := []string{"--", ";", "\"", " "}
	for _, s := range incorrectSymbols {
		cnt := strings.Count(value, s)
		if cnt != 0 {
			return false
		}
	}
	return true
}
