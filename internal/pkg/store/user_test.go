package store

import (
	"auth-service/cmd/configurations"
	"auth-service/internal/pkg/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_InsertUser(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		store := MustNew(configurations.StoreConfig)

		err := store.InsertUser(&domain.User{
			Username: "test_insert_user",
			Password: "test_insert_user",
		})

		require.NoError(t, err, "insert.user.error")

		_, err = store.db.Exec(`
		DELETE FROM users 
		WHERE username = 'test_insert_user' 
		AND PASSWORD = 'test_insert_user'
		`)
		require.NoError(t, err, "cleanup.error")
	})
	t.Run("dublicate", func(t *testing.T) {
		store := MustNew(configurations.StoreConfig)

		err := store.InsertUser(&domain.User{
			Username: "test_insert_user",
			Password: "test_insert_user",
		})
		require.NoError(t, err, "insert.user.error")

		err = store.InsertUser(&domain.User{
			Username: "test_insert_user",
			Password: "test_insert_user",
		})
		require.Error(t, err, "insert.dublicate.error")

		_, err = store.db.Exec(`
		DELETE FROM users 
		WHERE username = 'test_insert_user' 
		AND PASSWORD = 'test_insert_user'
		`)
		require.NoError(t, err, "cleanup.error")
	})
}
