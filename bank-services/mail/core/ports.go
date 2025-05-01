package core

type AuthorizationVerifier interface {
	// IdentifyUser going for Authorization service and request
	// identify user with username
	// return ErrUserNotExist if user not exist
	// return ErrUserWasBanned if user was banned
	IdentifyUser(username string) error
}

type SenderRecoverMessage interface {
	SendRecoverMessage(toUser string, code int) error
}

// there must be for generating reports
// I suppose it might get data from Report service
