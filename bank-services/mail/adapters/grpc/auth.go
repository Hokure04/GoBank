package grpc

import (
	authpb "github.com/Hokure04/GoBank/proto/auth"
	"google.golang.org/grpc"
	"log/slog"
)

type Auth struct {
	log    *slog.Logger
	client authpb.AuthClient
	conn   *grpc.ClientConn
}

func NewAuthClient(log *slog.Logger, conn *grpc.ClientConn) Auth {
	return Auth{
		log:    log,
		client: authpb.NewAuthClient(conn),
		conn:   conn,
	}
}


func