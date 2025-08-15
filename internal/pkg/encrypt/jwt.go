package encrypt

import (
	"auth-service/internal/pkg/config"

	"github.com/golang-jwt/jwt/v5"
)

type IJWT interface {
	iUser
}

type JWT struct {
	parser *jwt.Parser
	config *config.JWTConfig
}

func NewJWT(cfg *config.JWTConfig) *JWT {
	return &JWT{
		parser: jwt.NewParser(
			jwt.WithValidMethods([]string{"HS256"}),
			jwt.WithIssuedAt(),
			jwt.WithExpirationRequired(),
			jwt.WithLeeway(cfg.Leeway),
		),
		config: cfg,
	}
}
