package core

import "context"

type Pinger interface {
	Ping(context.Context) error
}

type GrpcClient interface {
	Close() error
	Pinger
}

type AuthorizationVerifier interface {
	GrpcClient
	// IdentifyUser going for Authorization service and request
	// identify user with username
	// return ErrUserNotExist if user not exist
	// return ErrUserWasBanned if user was banned
	IdentifyUser(ctx context.Context, username string) error
}

type SenderRecoverMessage interface {
	SendRecoverMessage(toUser string, code int) error
}

type Sender interface {
	SenderRecoverMessage
}

// there must be for generating reports
// I suppose it might get data from Report service
