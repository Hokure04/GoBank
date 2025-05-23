package grpcservice

import (
	"context"
	"github.com/Hokure04/GoBank/mail/core"
	authpb "github.com/Hokure04/GoBank/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (a Auth) RecoverPassword(ctx context.Context, email string) (string, error) {
	pass, err := a.client.RecoverPassword(ctx, &authpb.RecoverPass{
		Email: email,
	})

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return "", core.ErrUserNotExist
		}
		return "", err
	}
	return pass.GetPassword(), nil
}

func (a Auth) Close() error {
	if err := a.conn.Close(); err != nil {
		a.log.Error("ERROR while closing connection:", "error", err)
		return err
	}
	a.log.Debug("Words client are closed")
	return nil
}

func (a Auth) Ping(ctx context.Context) error {
	_, err := a.client.Ping(ctx, &emptypb.Empty{})
	return err
}
