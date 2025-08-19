package app

import (
	"auth-service/internal/pkg/mapper"
	"auth-service/internal/pkg/pb"
	"context"
)

type IUser interface {
	CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
}

func (s *Service) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	domainUser := mapper.User(in.GetUser())
	id, err := s.store.InsertUser(ctx, *domainUser)
	domainUser.ID = id
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{User: mapper.UserProto(domainUser)}, nil
}
