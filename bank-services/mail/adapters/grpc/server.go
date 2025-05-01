package grpc

import (
	"context"
	"errors"
	"github.com/Hokure04/GoBank/mail/core"
	mailpb "github.com/Hokure04/GoBank/proto/mail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

type Server struct {
	mailpb.UnimplementedMailServer
	log            *slog.Logger
	messageService core.MessageService
}

func NewGrpcServer(log *slog.Logger, messageService core.MessageService) Server {
	return Server{
		log:            log,
		messageService: messageService,
	}
}

func (s Server) Ping(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	s.log.Debug("Pinged")
	return nil, nil
}

func (s Server) RequestRecoverPassword(ctx context.Context, reg *mailpb.RecoverPass) (*emptypb.Empty, error) {
	if err := s.messageService.RecoverPassword(ctx, reg.GetEmail()); err != nil {
		if errors.Is(err, core.ErrUserNotExist) {
			s.log.Warn("user with such email not found", "username", reg.Email)
			return nil, status.Error(codes.NotFound, "user with such email not found")
		}
		s.log.Error("fail to send a recover message", "reason", err)
		return nil, status.Error(codes.Internal, "unpredictable error")
	}
	return nil, nil
}
