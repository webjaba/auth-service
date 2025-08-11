package jwttoken

import (
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	// setup
	config.MustLoadEnv("../../../.env")
	m := &JWTManager{config.LoadJWTConfig()}

	user := &domain.User{
		ID: 10,
	}

	// test
	tokenString, err := m.CreateFromUser(user)
	require.NoError(t, err, "CreateFromUser.error")

	token, err := m.ParseToken(tokenString)
	require.NoError(t, err, "ParseToken.error")
	require.NotNil(t, token, "ParseToken.token")

	verified, err := m.Validate(user, token)
	require.NoError(t, err, "Verify.error")
	require.True(t, verified, "Verify.verified")
}
