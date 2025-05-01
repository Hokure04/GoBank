package mail

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"log/slog"
)

var recoverTemplate *template.Template

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
	// TODO: добавить проверку на @... у пользователя

	return nil
}
