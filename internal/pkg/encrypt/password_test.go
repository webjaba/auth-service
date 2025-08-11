package encrypt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := "password"
	password2 := "password2"

	hash, err := EncryptPassword(password)
	require.NoError(t, err, "MustEncryptPassword.error")
	require.NotEmpty(t, hash, "MustEncryptPassword.hash")

	ok := IsCorrectPassword(password, hash)
	require.True(t, ok, "IsCorrectPassword")

	hash2, err := EncryptPassword(password2)
	require.NoError(t, err, "MustEncryptPassword2.error")
	require.NotEmpty(t, hash2, "MustEncryptPassword2.hash")

	ok = IsCorrectPassword(password, hash2)
	require.False(t, ok, "password - hash2")

	ok = IsCorrectPassword(password2, hash)
	require.False(t, ok, "password2 - hash")
}
