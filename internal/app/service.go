package app

import (
	"auth-service/internal/pkg/configuration"
	"auth-service/internal/pkg/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type IService interface {
	Start(configuration.ServerConfig) error
	Auth
}

type Service struct {
	pb.UnimplementedAuthServiceServer
	server *grpc.Server
}

func New() IService {
	server := grpc.NewServer()

	service := &Service{
		server: server,
	}

	pb.RegisterAuthServiceServer(server, service)
	return service
}

func (s Service) Start(cfg configuration.ServerConfig) error {
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
