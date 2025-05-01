package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Hokure04/GoBank/mail/adapters/grpcservice"
	"github.com/Hokure04/GoBank/mail/adapters/mail"
	"github.com/Hokure04/GoBank/mail/adapters/mail/stub"
	"github.com/Hokure04/GoBank/mail/config"
	"github.com/Hokure04/GoBank/mail/core"
	mailpb "github.com/Hokure04/GoBank/proto/mail"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"gopkg.in/gomail.v2"
	"os/signal"

	"log/slog"
	"net"
	"os"
)

const defaultConfigPath = "/Users/jaba/GolandProjects/extra/bank-services/pkg/config/mail/config.yaml"

const (
	mailAddressEnv  = "MAIL_ADDRESS"
	mailPasswordEnv = "MAIL_PASS"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", defaultConfigPath, "server configuration file")
	flag.Parse()
	cfg := config.MustLoad(configPath)

	log := mustMakeLogger(cfg.LogLevel)
	log.Info("logger is ready", "common config", cfg)
	log.Debug("debug level in logger is set")

	mailAddress, ok := os.LookupEnv(mailAddressEnv)
	if !ok {
		log.Error(fmt.Sprintf("%s must be set in env", mailAddressEnv))
		os.Exit(1)
	}
	mailPass, ok := os.LookupEnv(mailPasswordEnv)
	if !ok {
		log.Error(fmt.Sprintf("%s must be set in env", mailPasswordEnv))
		os.Exit(1)
	}

	sender := mustGetMailSender(log, mailCfg{
		appLevel:     cfg.AppLevel,
		username:     mailAddress,
		password:     mailPass,
		port:         465,
		host:         "smtp.mail.ru",
		pathTemplate: cfg.TemplatePath,
	})

	authConn, err := grpc.NewClient(cfg.Services.AuthServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("fail to connect to auth service")
		os.Exit(1)
	}
	authGrpcClient := grpcservice.NewAuthClient(log, authConn)
	messageService := core.NewMessageService(log, authGrpcClient, sender)

	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		log.Error("failed to listen", "error", err)
		os.Exit(1)
	}
	log.Info("server listening on", "address", cfg.Address)

	s := grpc.NewServer()
	mailpb.RegisterMailServer(s, grpcservice.NewGrpcServer(log, messageService))
	reflection.Register(s)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		<-ctx.Done()
		log.Debug("shutting down server")
		s.GracefulStop()
	}()

	if err := s.Serve(listener); err != nil {
		log.Error("failed to serve", "error", err)
	}

}

func mustMakeLogger(logLevel string) *slog.Logger {
	var level slog.Level
	switch logLevel {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "ERROR":
		level = slog.LevelError
	default:
		panic("unknown log level: " + logLevel)
	}
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	return slog.New(handler)
}

type mailCfg struct {
	appLevel     string
	username     string
	password     string
	port         int
	host         string
	pathTemplate string
}

func mustGetMailSender(log *slog.Logger, cfg mailCfg) core.Sender {
	switch cfg.appLevel {
	case "PROD":
		d := gomail.NewDialer(cfg.host, cfg.port, cfg.username, cfg.password)
		return mail.NewMailSender(log, d, cfg.pathTemplate)
	default:
		return stub.NewStubMailSender(log, cfg.username, cfg.pathTemplate)
	}
}
