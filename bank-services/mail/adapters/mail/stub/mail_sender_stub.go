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
	"time"
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
	dialer string
}

func NewStubMailSender(log *slog.Logger, dialer string, templateFolder string) Sender {
	recoverTemplate = template.Must(
		template.ParseFiles(
			fmt.Sprintf("%s/recovermessage.gohtml", templateFolder)))
	return Sender{
		log:    log,
		dialer: dialer,
	}
}

func (s Sender) SendRecoverMessage(toUser string, pass string) error {
	if !validateEmail(toUser) {
		return ErrInvalidUser
	}

	var body bytes.Buffer

	err := recoverTemplate.Execute(&body, struct {
		Pass string
	}{
		Pass: pass,
	})
	if err != nil {
		return err
	}

	message := s.generateMessage(toUser, body.String(), headerResetPassword)
	s.log.Info("Sending to user data...", "message", message)
	time.Sleep(time.Second * 2)

	return nil
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (s Sender) generateMessage(email, body, header string) *gomail.Message {
	message := gomail.NewMessage()

	message.SetHeader("From", s.dialer)
	message.SetHeader("To", email)
	message.SetHeader("Subject", header)
	message.SetBody("text/html", body)

	return message
}
