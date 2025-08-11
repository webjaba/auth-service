package store

import (
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/domain"
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
	id, err := store.InsertUser(user)
	require.NoError(t, err, "InsertUser.error.first")

	created, err := store.GetUserByID(id)
	require.NoError(t, err, "GetUserByID.error")
	require.Equal(t, created.Username, user.Username, "GetUserByID.username")

	_, err = store.InsertUser(user)
	require.Error(t, err, "InsertUser.error.second")

}

func cleanup(t *testing.T, store *Store) {
	_, cleanupError := store.db.Exec(`
	DELETE FROM users 
	WHERE username = 'test_insert_user'
	`)
	require.NoError(t, cleanupError, "cleanup_error")
}
