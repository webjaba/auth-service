package encrypt

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type iUser interface {
	Parse(raw string) (*jwt.Token, error)
	Create(id int) (string, error)
}

func (j *JWT) Parse(raw string) (*jwt.Token, error) {
	token, err := j.parser.Parse(raw, func(t *jwt.Token) (any, error) {
		return j.config.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (j *JWT) Create(id int) (string, error) {
	iat := time.Now()
	exp := iat.Add(j.config.Duration)
	claims := jwt.MapClaims{
		"sub": strconv.Itoa(id),
		"exp": jwt.NewNumericDate(exp),
		"iat": jwt.NewNumericDate(iat),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(j.config.SecretKey)
	if err != nil {
		return "", err
	}
	return signed, nil
}
