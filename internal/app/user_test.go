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

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(s *Setup) *pb.CreateUserResponse
		in      *pb.CreateUserRequest
		wantErr bool
	}{
		{
			name: "ok",
			setup: func(s *Setup) *pb.CreateUserResponse {
				s.store.
					EXPECT().
					InsertUser(domain.User{
						Username: "user1",
						Password: "password1",
					}).Return(1, nil)
				return &pb.CreateUserResponse{
					User: &pb.User{
						Id:       1,
						Username: "user1",
						Password: "password1",
					},
				}
			},
			in: &pb.CreateUserRequest{
				User: &pb.User{
					Username: "user1",
					Password: "password1",
				},
			},
			wantErr: false,
		},
		{
			name: "insert error",
			setup: func(s *Setup) *pb.CreateUserResponse {
				s.store.
					EXPECT().
					InsertUser(domain.User{
						Username: "user1",
						Password: "password1",
					}).Return(0, errors.New("err"))
				return nil
			},
			in: &pb.CreateUserRequest{
				User: &pb.User{
					Username: "user1",
					Password: "password1",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			s := MustSetup(ctrl)

			want := tt.setup(s)

			got, err := s.CreateUser(context.Background(), tt.in)
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, got)
				return
			}
			require.NoError(t, err)
			require.Equal(t, want.User.Id, got.User.Id, "User.Id")
			require.Equal(t, want.User.Username, got.User.Username, "User.Username")
			require.Equal(t, want.User.Password, got.User.Password, "User.Password")
		})
	}
}

func TestCreateToken(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(s *Setup) *pb.CreateTokenResponse
		in      *pb.CreateTokenRequest
		wantErr bool
	}{
		{
			name: "ok",
			setup: func(s *Setup) *pb.CreateTokenResponse {
				in := mustBuildCreateTokenRequest(1, "user", "password")

				password, _ := encrypt.EncryptPassword(in.User.Password)

				domainUser := &domain.User{
					ID:       int(in.User.Id),
					Username: in.User.Username,
					Password: password,
				}

				s.store.EXPECT().GetUserByID(1).
					Return(domainUser, nil)

				s.jwtManager.EXPECT().
					CreateFromUser(domainUser).
					Return("token", nil)

				return &pb.CreateTokenResponse{Token: "token"}
			},
			in:      mustBuildCreateTokenRequest(1, "user", "password"),
			wantErr: false,
		},
		{
			name: "error user does not exists",
			setup: func(s *Setup) *pb.CreateTokenResponse {
				domainUser := &domain.User{
					ID:       1,
					Username: "",
					Password: "",
				}

				s.store.EXPECT().GetUserByID(1).
					Return(domainUser, nil)

				return nil
			},
			in:      mustBuildCreateTokenRequest(1, "", ""),
			wantErr: true,
		},
		{
			name: "error incorrect username",
			setup: func(s *Setup) *pb.CreateTokenResponse {
				in := mustBuildCreateTokenRequest(1, "user1", "")

				domainUser := &domain.User{
					ID:       int(in.User.Id),
					Username: "user2",
					Password: "",
				}

				s.store.EXPECT().GetUserByID(1).
					Return(domainUser, nil)

				return nil
			},
			in:      mustBuildCreateTokenRequest(1, "user1", ""),
			wantErr: true,
		},
		{
			name: "error incorrect password",
			setup: func(s *Setup) *pb.CreateTokenResponse {
				in := mustBuildCreateTokenRequest(1, "user1", "password1")

				p, _ := encrypt.EncryptPassword("password2")
				domainUser := &domain.User{
					ID:       int(in.User.Id),
					Username: "user1",
					Password: p,
				}

				s.store.EXPECT().GetUserByID(1).
					Return(domainUser, nil)

				return nil
			},
			in:      mustBuildCreateTokenRequest(1, "user1", "password1"),
			wantErr: true,
		},
		{
			name: "store error",
			setup: func(s *Setup) *pb.CreateTokenResponse {
				in := mustBuildCreateTokenRequest(1, "user", "password")

				p, _ := encrypt.EncryptPassword(in.User.Password)
				domainUser := &domain.User{
					ID:       int(in.User.Id),
					Username: in.User.Username,
					Password: p,
				}

				s.store.EXPECT().GetUserByID(1).
					Return(domainUser, errors.New("error"))

				return nil
			},
			in:      mustBuildCreateTokenRequest(1, "user", "password"),
			wantErr: true,
		},
		{
			name: "jwt error",
			setup: func(s *Setup) *pb.CreateTokenResponse {
				in := mustBuildCreateTokenRequest(1, "user", "password")

				p, _ := encrypt.EncryptPassword(in.User.Password)
				domainUser := &domain.User{
					ID:       int(in.User.Id),
					Username: in.User.Username,
					Password: p,
				}

				s.store.EXPECT().GetUserByID(1).
					Return(domainUser, nil)

				s.jwtManager.EXPECT().
					CreateFromUser(domainUser).
					Return("", errors.New("error"))

				return nil
			},
			in:      mustBuildCreateTokenRequest(1, "user", "password"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			s := MustSetup(ctrl)

			want := tt.setup(s)

			got, err := s.CreateToken(context.Background(), tt.in)
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, got)
				return
			}
			require.NoError(t, err)
			require.Equal(t, want.Token, got.Token)
		})
	}
}

func mustBuildCreateTokenRequest(id int64, username, password string) *pb.CreateTokenRequest {
	return &pb.CreateTokenRequest{
		User: &pb.User{
			Id:       id,
			Username: username,
			Password: password,
		},
	}
}
