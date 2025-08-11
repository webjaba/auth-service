package app

import (
	"auth-service/internal/pkg/encrypt"
	"auth-service/internal/pkg/mapper"
	"auth-service/internal/pkg/pb"
	"context"
	"errors"
)

var (
	errIcorrectPassword  = errors.New("incorrect password")
	errIcorrectUsername  = errors.New("incorrect username")
	errUserDoesNotExists = errors.New("user does not exists")
	errInvalidToken      = errors.New("invalid token")
)

type Auth interface {
	CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	CreateToken(ctx context.Context, in *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error)
	RecreateToken(ctx context.Context, in *pb.RecreateTokenRequest) (*pb.RecreateTokenResponse, error)
	VerifyToken(ctx context.Context, in *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error)
}

func (s *Service) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	domainUser := mapper.User(in.GetUser())
	id, err := s.store.InsertUser(*domainUser)
	domainUser.ID = id
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{User: mapper.UserProto(domainUser)}, nil
}

func (s *Service) CreateToken(ctx context.Context, in *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
	domainUser, err := s.store.GetUserByID(int(in.User.Id))
	if err != nil {
		return nil, err
	}
	if domainUser.Username == "" {
		return nil, errUserDoesNotExists
	}
	if domainUser.Username != in.User.Username {
		return nil, errIcorrectUsername
	}
	if ok := encrypt.IsCorrectPassword(in.User.Password, domainUser.Password); !ok {
		return nil, errIcorrectPassword
	}

	token, err := s.jwtManager.CreateFromUser(domainUser)
	if err != nil {
		return nil, err
	}
	return &pb.CreateTokenResponse{Token: token}, nil
}

func (s *Service) RecreateToken(ctx context.Context, in *pb.RecreateTokenRequest) (*pb.RecreateTokenResponse, error) {
	return nil, nil
}

func (s *Service) VerifyToken(ctx context.Context, in *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	return nil, nil
}
