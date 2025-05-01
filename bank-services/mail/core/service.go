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
	pass, err := m.auth.RecoverPassword(ctx, username)
	if err != nil {
		if errors.Is(err, ErrUserNotExist) {
			m.log.Warn("user with such username does not exist", "reason", err)
			return err
		} else {
			m.log.Error("unpredictable error", "reason", err)
			return err
		}
	}
	if err := m.sender.SendRecoverMessage(username, pass); err != nil {
		// TODO: сюда ИЛИ распределённую транзакцию ИЛИ САГУ, потому что пароль уже сброшен,
		// 	а пользователь не оповещен
		m.log.Error("fail to send a message")
		return err
	}
	return nil
}
