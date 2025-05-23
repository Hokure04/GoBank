package mail

import (
	"gopkg.in/gomail.v2"
	"log/slog"
	"testing"
)

func TestSender_SendRecoverMessage(t *testing.T) {
	validDieler := gomail.NewDialer("smtp.mail.ru", 465, "mail", "password")

	type fields struct {
		log    *slog.Logger
		dialer *gomail.Dialer
	}
	type args struct {
		toUser string
		code   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "full validateEmail test",
			fields: fields{
				log:    slog.Default(),
				dialer: validDieler,
			},
			args: args{
				toUser: "some@gmail.com",
				code:   "3421",
			},
			wantErr: false,
		},
		{
			name: "invalid email",
			fields: fields{
				log:    slog.Default(),
				dialer: validDieler,
			},
			args: args{
				toUser: "jaba",
				code:   "21",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewMailSender(tt.fields.log, tt.fields.dialer, "/Users/jaba/GolandProjects/extra/bank-services/mail/templates")
			if err := s.SendRecoverMessage(tt.args.toUser, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("SendRecoverMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
