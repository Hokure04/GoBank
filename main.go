package main

import (
	"PetGoProject/models/user"
	"PetGoProject/sever"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Ошибка запуска сервера")
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &sever.UserServiceServer{})

	log.Println("grpc сервер запущен")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
