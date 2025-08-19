package store

import (
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/domain"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUser(t *testing.T) {
	// setup
	config.MustLoadEnv("../../../.env")
	store := MustNew(config.LoadStoreConfig())
	user := domain.User{
		Username: "test_insert_user",
		Password: "test_insert_user",
	}
	// cleanup
	defer cleanup(t, store)

	// test
	_, err := store.InsertUser(context.Background(), user)
	require.NoError(t, err, "InsertUser.error.first")

	created, err := store.FindUser(context.Background(), "test_insert_user")
	require.NoError(t, err, "FindUser.error")
	require.Equal(t, created.Username, user.Username, "FindUser.username")

	_, err = store.InsertUser(context.Background(), user)
	require.Error(t, err, "InsertUser.error.second")

}

func cleanup(t *testing.T, store *Store) {
	_, cleanupError := store.db.Exec(`
	DELETE FROM users 
	WHERE username = 'test_insert_user'
	`)
	require.NoError(t, cleanupError, "cleanup_error")
}
