package app

import (
	"auth-service/internal/pkg/encrypt"
	"auth-service/internal/pkg/pb"
	"context"
	"errors"
)

var (
	errIcorrectPassword = errors.New("incorrect password")
)

type IJWT interface {
	CreateToken(ctx context.Context, in *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error)
	RecreateToken(ctx context.Context, in *pb.RecreateTokenRequest) (*pb.RecreateTokenResponse, error)
}

func (s *Service) CreateToken(ctx context.Context, in *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
	domainUser, err := s.store.FindUser(in.User.Username)
	if err != nil {
		return nil, err
	}

	ok := encrypt.IsCorrectPassword(in.User.Password, domainUser.Password)
	if !ok {
		return nil, errIcorrectPassword
	}

	token, err := s.jwt.Create(domainUser.ID)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTokenResponse{Token: token}, nil
}

func (s *Service) RecreateToken(ctx context.Context, in *pb.RecreateTokenRequest) (*pb.RecreateTokenResponse, error) {
	id, err := s.jwt.ParseId(in.Token)
	if err != nil {
		return nil, err
	}

	tokenString, err := s.jwt.Create(id)
	if err != nil {
		return nil, err
	}

	return &pb.RecreateTokenResponse{Token: tokenString}, nil
}
