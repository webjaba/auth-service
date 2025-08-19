package app

import (
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/encrypt"
	"auth-service/internal/pkg/pb"
	"auth-service/internal/pkg/store"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type IService interface {
	Start(*config.ServerConfig) error
	IUser
	IJWT
}

type Service struct {
	pb.UnimplementedAuthServiceServer
	server *grpc.Server
	store  store.IStore
	jwt    encrypt.IJWT
}

func New(server *grpc.Server, store store.IStore, jwtManager encrypt.IJWT) IService {
	service := &Service{
		server: server,
		store:  store,
		jwt:    jwtManager,
	}

	pb.RegisterAuthServiceServer(server, service)
	return service
}

func (s Service) Start(cfg *config.ServerConfig) error {
	reflection.Register(s.server)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if err != nil {
		return err
	}

	if err := s.server.Serve(lis); err != nil {
		return err
	}

	return nil
}
