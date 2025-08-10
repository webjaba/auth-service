package store

import (
	"auth-service/cmd/configurations"
	"auth-service/internal/pkg/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_InsertUser(t *testing.T) {
	// setup
	store := MustNew(configurations.StoreConfig)
	var (
		insertUserError error
		dublicateError  error
		cleanupError    error
	)

	// test
	_, insertUserError = store.InsertUser(domain.User{
		Username: "test_insert_user",
		Password: "test_insert_user",
	})

	_, dublicateError = store.InsertUser(domain.User{
		Username: "test_insert_user",
		Password: "test_insert_user",
	})

	// cleanup
	_, cleanupError = store.db.Exec(`
	DELETE FROM users 
	WHERE username = 'test_insert_user'
	`)

	// error check
	require.NoError(t, insertUserError, "insert_user_error")
	require.Error(t, dublicateError, "insert_dublicate_error")
	require.NoError(t, cleanupError, "cleanup_error")
}
