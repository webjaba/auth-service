package store

import (
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_InsertUser(t *testing.T) {
	// setup
	config.MustLoadEnv("../../../.env")
	store := MustNew(config.LoadStoreConfig())
	var (
		insertUserError error
		dublicateError  error
	)

	// cleanup
	defer cleanup(t, store)

	// test
	_, insertUserError = store.InsertUser(domain.User{
		Username: "test_insert_user",
		Password: "test_insert_user",
	})

	_, dublicateError = store.InsertUser(domain.User{
		Username: "test_insert_user",
		Password: "test_insert_user",
	})

	// error check
	require.NoError(t, insertUserError, "insert_user_error")
	require.Error(t, dublicateError, "insert_dublicate_error")
}

func cleanup(t *testing.T, store *Store) {
	_, cleanupError := store.db.Exec(`
	DELETE FROM users 
	WHERE username = 'test_insert_user'
	`)
	require.NoError(t, cleanupError, "cleanup_error")
}
