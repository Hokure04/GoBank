package main

import (
	"PetGoProject/models/user"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Ошибка подключения", err)
	}

	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	ctx, cncel := context.WithTimeout(context.Background(), time.Second)
	defer cncel()

	result, err := client.RegisterUser(ctx, &user.RegisterUserRequest{
		FullName: "Каргин Александр",
		Email:    "some@mail.ru",
		Password: "12345678",
	})
	if err != nil {
		log.Fatal("Ошибка при регистрации пользователя", err)
	}

	log.Printf("Данные от сервера: %s, %s", result.Message, result.UserId)
}
