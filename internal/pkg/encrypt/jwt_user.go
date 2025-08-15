package encrypt

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type iUser interface {
	ParseId(raw string) (int, error)
	Create(id int) (string, error)
}

func (j *JWT) parse(raw string) (*jwt.Token, error) {
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

func (j *JWT) ParseId(raw string) (int, error) {
	token, err := j.parse(raw)
	if err != nil {
		return 0, err
	}

	sub, err := token.Claims.GetSubject()
	if err != nil {
		return 0, err
	}

	id, err := strconv.Atoi(sub)
	if err != nil {
		return 0, err
	}

	return id, nil
}
