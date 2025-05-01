package core

import "errors"

var (
	ErrUserNotExist  = errors.New("user with this username not found")
	ErrUserWasBanned = errors.New("user with this username was banned")
)
