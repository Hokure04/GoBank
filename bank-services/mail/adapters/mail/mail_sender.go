// mail service for processing real work with mails
// if you work in dev mode use mail_sender_stub.go -> it does the same thing, but without sending mails

package mail

import (
	"bytes"
	"errors"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"log/slog"
	"net/mail"
)

var recoverTemplate *template.Template

var (
	ErrInvalidUser = errors.New("invalid user")
)

const (
	headerResetPassword = "Reset password"
)

type Sender struct {
	log    *slog.Logger
	dialer *gomail.Dialer
}

func NewMailSender(log *slog.Logger, dialer *gomail.Dialer, templateFolder string) Sender {
	recoverTemplate = template.Must(
		template.ParseFiles(
			fmt.Sprintf("%s/recovermessage.gohtml", templateFolder)))
	return Sender{
		log:    log,
		dialer: dialer,
	}
}

func (s Sender) SendRecoverMessage(toUser string, code int) error {
	if !validateEmail(toUser) {
		s.log.Warn("user email is invalid", "email", toUser)
		return ErrInvalidUser
	}

	var body bytes.Buffer

	err := recoverTemplate.Execute(&body, struct {
		Code int
	}{
		Code: code,
	})
	if err != nil {
		s.log.Error("fail to execute data for template")
		return err
	}

	message := s.generateMessage(toUser, body.String(), headerResetPassword)
	if err = s.dialer.DialAndSend(message); err != nil {
		s.log.Error("fail to send a message", "reason", err)
		return err
	}

	return nil
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (s Sender) generateMessage(email, body, header string) *gomail.Message {
	message := gomail.NewMessage()

	message.SetHeader("From", s.dialer.Username)
	message.SetHeader("To", email)
	message.SetHeader("Subject", header)
	message.SetBody("text/html", body)

	return message
}
