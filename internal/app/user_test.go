package app

import (
	"auth-service/internal/pkg/domain"
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
					InsertUser(
						gomock.Any(),
						domain.User{
							Username: "user1",
							Password: "password1",
						},
					).Return(1, nil)
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
					InsertUser(
						gomock.Any(),
						domain.User{
							Username: "user1",
							Password: "password1",
						},
					).Return(0, errors.New("err"))
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
