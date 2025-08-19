package store

import (
	"auth-service/internal/pkg/domain"
	"auth-service/internal/pkg/encrypt"
	"context"
	"errors"
	"strings"
)

var (
	errInvalidUsername = errors.New("invalid username")
)

type IUser interface {
	InsertUser(ctx context.Context, user domain.User) (id int, err error)
	FindUser(ctx context.Context, username string) (*domain.User, error)
}

func (s *Store) InsertUser(ctx context.Context, user domain.User) (id int, err error) {
	if ok := validateString(user.Username); !ok {
		return 0, errInvalidUsername
	}

	user.Password, err = encrypt.EncryptPassword(user.Password)
	if err != nil {
		return -1, err
	}

	err = s.db.GetContext(
		ctx, &id,
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id",
		user.Username, user.Password,
	)
	if err != nil {
		return -1, err
	}

	return
}

func (s *Store) FindUser(ctx context.Context, username string) (*domain.User, error) {
	if !validateString(username) {
		return nil, errInvalidUsername
	}
	user := &domain.User{}
	err := s.db.GetContext(ctx, user, "SELECT id, username, password FROM users WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func validateString(value string) bool {
	incorrectSymbols := []string{"--", ";", "\"", " "}
	for _, s := range incorrectSymbols {
		cnt := strings.Count(value, s)
		if cnt != 0 {
			return false
		}
	}
	return true
}
