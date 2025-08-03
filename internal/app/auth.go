package app

import (
	"auth-service/internal/pkg/pb"
	"context"
)

type Auth interface {
	CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
}

func (s *Service) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return nil, nil
}
