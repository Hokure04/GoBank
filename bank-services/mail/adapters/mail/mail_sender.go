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
		return ErrInvalidUser
	}

	var body bytes.Buffer

	err := recoverTemplate.Execute(&body, struct {
		Code int
	}{
		Code: code,
	})
	if err != nil {
		return err
	}

	// TODO возможно позже вынести в отдельную функцию
	message := gomail.NewMessage()

	message.SetHeader("From", s.dialer.Username)
	message.SetHeader("To", toUser)
	message.SetHeader("Subject", "Reset password")
	message.SetBody("text/html", body.String())

	if err = s.dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
