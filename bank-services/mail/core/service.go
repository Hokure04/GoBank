// Service which manipulate messages(for now only email)
package core

import (
	"context"
	"errors"
	"log/slog"
)

type MessageService struct {
	log    *slog.Logger
	auth   AuthorizationVerifier
	sender Sender
}

func NewMessageService(log *slog.Logger, auth AuthorizationVerifier, sender Sender) MessageService {
	return MessageService{
		log:    log,
		auth:   auth,
		sender: sender,
	}
}

func (m MessageService) RecoverPassword(ctx context.Context, username string) error {
	if err := m.auth.IdentifyUser(ctx, username); err != nil {
		if errors.Is(err, ErrUserNotExist) {
			m.log.Warn("user with such username does not exist", "reason", err)
			return err
		} else {
			m.log.Error("unpredictable error", "reason", err)
			return err
		}
	}

	if err := m.sender.SendRecoverMessage(username); err != nil {

	}
}
