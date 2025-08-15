package app

import (
	"auth-service/internal/pkg/domain"
	"auth-service/internal/pkg/encrypt"
	"auth-service/internal/pkg/pb"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_CreateToken(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(t *testing.T, s *Setup) (req *pb.CreateTokenRequest)
		wantErr bool
	}{
		{
			name: "ok",
			setup: func(t *testing.T, s *Setup) (req *pb.CreateTokenRequest) {
				req = &pb.CreateTokenRequest{
					User: &pb.User{
						Username: "user",
						Password: "password",
					},
				}

				p, err := encrypt.EncryptPassword(req.User.Password)
				require.NoError(t, err, "hash password error")

				s.store.EXPECT().
					FindUser(req.User.Username).
					Return(&domain.User{
						ID:       1,
						Username: req.User.Username,
						Password: p,
					}, nil)

				s.jwt.EXPECT().Create(1).Return("token", nil)

				return
			},
			wantErr: false,
		},
		{
			name: "find user error",
			setup: func(t *testing.T, s *Setup) (req *pb.CreateTokenRequest) {
				req = &pb.CreateTokenRequest{
					User: &pb.User{
						Username: "user",
						Password: "password",
					},
				}

				s.store.EXPECT().
					FindUser(req.User.Username).
					Return(nil, errors.New("error"))

				return
			},
			wantErr: true,
		},
		{
			name: "incorrect password",
			setup: func(t *testing.T, s *Setup) (req *pb.CreateTokenRequest) {
				req = &pb.CreateTokenRequest{
					User: &pb.User{
						Username: "user",
						Password: "password",
					},
				}

				p, err := encrypt.EncryptPassword("another_password")
				require.NoError(t, err, "hash password error")

				s.store.EXPECT().
					FindUser(req.User.Username).
					Return(&domain.User{
						ID:       1,
						Username: req.User.Username,
						Password: p,
					}, nil)

				return
			},
			wantErr: true,
		},
		{
			name: "err create token",
			setup: func(t *testing.T, s *Setup) (req *pb.CreateTokenRequest) {
				req = &pb.CreateTokenRequest{
					User: &pb.User{
						Username: "user",
						Password: "password",
					},
				}

				p, err := encrypt.EncryptPassword(req.User.Password)
				require.NoError(t, err, "hash password error")

				s.store.EXPECT().
					FindUser(req.User.Username).
					Return(&domain.User{
						ID:       1,
						Username: req.User.Username,
						Password: p,
					}, nil)

				s.jwt.EXPECT().Create(1).Return("", errors.New("error"))

				return
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := MustSetup(ctrl)

			in := tt.setup(t, s)
			got, err := s.CreateToken(context.Background(), in)
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, got)
				return
			}
			require.NoError(t, err)
			require.NotEmpty(t, got.Token)
		})
	}
}

func Test_RecreateToken(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(s *Setup) (req *pb.RecreateTokenRequest)
		wantErr bool
	}{
		{
			name: "ok",
			setup: func(s *Setup) (req *pb.RecreateTokenRequest) {
				req = &pb.RecreateTokenRequest{
					Token: "token",
				}
				s.jwt.EXPECT().ParseId("token").Return(1, nil)

				s.jwt.EXPECT().Create(1).Return("token", nil)

				return
			},
			wantErr: false,
		},
		{
			name: "parsing error",
			setup: func(s *Setup) (req *pb.RecreateTokenRequest) {
				req = &pb.RecreateTokenRequest{
					Token: "token",
				}
				s.jwt.EXPECT().ParseId("token").Return(0, errors.New("error"))

				return
			},
			wantErr: true,
		},
		{
			name: "creation error",
			setup: func(s *Setup) (req *pb.RecreateTokenRequest) {
				req = &pb.RecreateTokenRequest{
					Token: "token",
				}
				s.jwt.EXPECT().ParseId("token").Return(1, nil)

				s.jwt.EXPECT().Create(1).Return("", errors.New("error"))

				return
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := MustSetup(ctrl)

			in := tt.setup(s)
			got, err := s.RecreateToken(context.Background(), in)
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, got)
				return
			}
			require.NoError(t, err)
			require.NotEmpty(t, got.Token)
		})
	}
}
