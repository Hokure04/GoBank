package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func main() {
	// Create a new message
	message := gomail.NewMessage()

	// Set email headers
	message.SetHeader("From", "go.bank.03@mail.ru")
	message.SetHeader("To", "random@gmail.com")
	message.SetHeader("Subject", "Hello from the Mailtrap team")

	// Set email body
	message.SetBody("text/plain", "This is the Test Body")

	// Set up the SMTP dialer Vjrm6i0L46zzKntnEEkD
	dialer := gomail.NewDialer("smtp.mail.ru", 465, "go.bank.03@mail.ru", "Vjrm6i0L46zzKntnEEkD")

	//// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		panic(err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
