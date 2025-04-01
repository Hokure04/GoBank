package sever

import (
	"PetGoProject/models/user"
	"context"
	"log"
)

type UserServiceServer struct {
	user.UnimplementedUserServiceServer
}

func (s *UserServiceServer) RegisterUser(ctx context.Context, request *user.RegisterUserRequest) (*user.RegisterUserResponse, error) {
	log.Printf("НПользователь: %s", request.FullName)

	return &user.RegisterUserResponse{
		UserId:  "1",
		Message: "Пользователь зарегистрирован",
	}, nil
}
