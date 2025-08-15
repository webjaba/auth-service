package encrypt

import (
	"auth-service/internal/pkg/config"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func Test_Parse(t *testing.T) {
	tests := []struct {
		name     string
		setupJWT func() IJWT
		setup    func() string
		want     string
		wantErr  bool
	}{
		{
			name: "expired",
			setup: func() string {
				iat := time.Now()
				exp := iat.Add(100 * time.Minute)
				claims := jwt.MapClaims{
					"sub": "1",
					"exp": jwt.NewNumericDate(exp),
					"iat": jwt.NewNumericDate(iat),
				}

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

				signed, _ := token.SignedString([]byte("key"))
				return signed
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "expired",
			setup: func() string {
				iat := time.Now()
				claims := jwt.MapClaims{
					"sub": "",
					"exp": jwt.NewNumericDate(iat),
					"iat": jwt.NewNumericDate(iat),
				}

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

				signed, _ := token.SignedString([]byte("key"))

				time.Sleep(10 * time.Millisecond)
				return signed
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "without exp",
			setup: func() string {
				iat := time.Now()
				claims := jwt.MapClaims{
					"sub": "",
					"iat": jwt.NewNumericDate(iat),
				}

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

				signed, _ := token.SignedString([]byte("key"))

				return signed
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "without iat",
			setup: func() string {
				exp := time.Now()
				claims := jwt.MapClaims{
					"sub": "",
					"exp": jwt.NewNumericDate(exp),
				}

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

				signed, _ := token.SignedString([]byte("key"))

				return signed
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "different signing method",
			setup: func() string {
				iat := time.Now()
				exp := iat.Add(100 * time.Minute)
				claims := jwt.MapClaims{
					"sub": "1",
					"exp": jwt.NewNumericDate(exp),
					"iat": jwt.NewNumericDate(iat),
				}

				token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

				signed, _ := token.SignedString([]byte("key"))
				return signed
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "different key",
			setup: func() string {
				iat := time.Now()
				exp := iat.Add(100 * time.Minute)
				claims := jwt.MapClaims{
					"sub": "1",
					"exp": jwt.NewNumericDate(exp),
					"iat": jwt.NewNumericDate(iat),
				}

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

				signed, _ := token.SignedString([]byte("different_key"))
				return signed
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := NewJWT(&config.JWTConfig{
				SecretKey: []byte("key"),
			})

			token, err := j.Parse(tt.setup())
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, token)
				return
			}
			require.NoError(t, err)

			got, _ := token.Claims.GetSubject()
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_Create(t *testing.T) {
	j := NewJWT(&config.JWTConfig{
		Duration:  10 * time.Minute,
		SecretKey: []byte("key"),
	})

	token, err := j.Create(1)
	require.NoError(t, err)

	parsed, err := j.Parse(token)
	require.NoError(t, err)

	sub, err := parsed.Claims.GetSubject()
	require.NoError(t, err)

	require.Equal(t, sub, "1")
}
