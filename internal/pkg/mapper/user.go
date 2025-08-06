package mapper

import (
	"auth-service/internal/pkg/domain"
	"auth-service/internal/pkg/pb"
)

func User(in *pb.User) *domain.User {
	return &domain.User{
		Username: in.GetUsername(),
		Password: in.GetPassword(),
	}
}

func UserProto(in *domain.User) *pb.User {
	return &pb.User{
		Id:       int64(in.ID),
		Username: in.Username,
		Password: in.Password,
	}
}
