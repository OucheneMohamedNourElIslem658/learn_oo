package email

import (
	"net/smtp"
)

var (
	auth smtp.Auth
)

func Init() {
	auth = smtp.PlainAuth(
		"",
		mailerConfig.email,
		mailerConfig.password,
		"smtp.gmail.com",
	)
}

func SendEmailMessage(toEmail string, message string) error {
	// Sending Email
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		mailerConfig.email,
		[]string{toEmail},
		[]byte(message),
	)

	return err
}
