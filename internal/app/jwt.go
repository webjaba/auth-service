package app

import (
	"auth-service/internal/pkg/pb"
	"context"
)

// var (
// 	errIcorrectPassword  = errors.New("incorrect password")
// 	errIcorrectUsername  = errors.New("incorrect username")
// 	errUserDoesNotExists = errors.New("user does not exists")
// 	errInvalidToken      = errors.New("invalid token")
// )

type IJWT interface {
	CreateToken(ctx context.Context, in *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error)
	RecreateToken(ctx context.Context, in *pb.RecreateTokenRequest) (*pb.RecreateTokenResponse, error)
}

// здесь надо получать id пользователя по username+password
// процесс является авторизацией
func (s *Service) CreateToken(ctx context.Context, in *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
	return nil, nil
}

// TODO: дописать 2 этих метода и написать на них тесты

func (s *Service) RecreateToken(ctx context.Context, in *pb.RecreateTokenRequest) (*pb.RecreateTokenResponse, error) {
	return nil, nil
}
