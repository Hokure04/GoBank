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
	// RecoverPassword going for Authorization service and request a temporary password for user
	// identify user with such email
	// return temporary password, which was already set in database
	// return ErrUserNotExist if user not exist
	// return ErrUserWasBanned if user was banned
	RecoverPassword(ctx context.Context, username string) (string, error)
}

type SenderRecoverMessage interface {
	SendRecoverMessage(toUser string, tmpPass string) error
}

type Sender interface {
	SenderRecoverMessage
}

// there must be for generating reports
// I suppose it might get data from Report service
