package encrypt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := "password"

	hash, err := EncryptPassword(password)
	require.NoError(t, err, "MustEncryptPassword.error")
	require.NotEmpty(t, hash, "MustEncryptPassword.hash")

	ok := IsCorrectPassword(password, hash)
	require.True(t, ok, "IsCorrectPassword")
}
