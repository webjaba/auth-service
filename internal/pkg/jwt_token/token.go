package jwttoken

import (
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/domain"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	errUnknownSigningMethod = errors.New("unknown signing method")
	errTokenExpired         = errors.New("expired token")
)

type JWTManager struct {
	*config.JWTConfig
}

func NewManager(cfg *config.JWTConfig) *JWTManager {
	return &JWTManager{cfg}
}

func (m *JWTManager) CreateFromUser(user *domain.User) (string, error) {
	iat := time.Now()
	exp := iat.Add(m.DurationMin)
	claims := jwt.MapClaims{
		"sub": fmt.Sprintf("%d", user.ID),
		"exp": jwt.NewNumericDate(exp),
		"iat": jwt.NewNumericDate(iat),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(m.SecretKey)
	if err != nil {
		return "", err
	}
	return signed, nil
}

func (m *JWTManager) Verify(user *domain.User, tokenString string) (bool, error) {
	token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errUnknownSigningMethod
		}
		return m.SecretKey, nil
	})

	sub, err := token.Claims.GetSubject()
	if err != nil {
		return false, err
	}
	iat, err := token.Claims.GetIssuedAt()
	if err != nil || iat == nil {
		return false, err
	}
	exp, err := token.Claims.GetExpirationTime()
	if err != nil || exp == nil {
		return false, err
	}

	if exp.Time.Unix() < time.Now().Unix() || exp.Time.Sub(iat.Time).Minutes() >= float64(m.DurationMin) {
		return false, errTokenExpired
	}

	return sub == fmt.Sprintf("%d", user.ID), nil
}
