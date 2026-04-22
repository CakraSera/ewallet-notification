package external

import (
	"strconv"

	"ewallet-notification/helper"

	"gopkg.in/gomail.v2"
)

type Email struct {
	To      string
	From    string
	Subject string
	Body    string
}

func (e *Email) SendEmail() error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("To", e.To)
	mailer.SetHeader("From", e.From)
	mailer.SetHeader("Subject", e.Subject)
	mailer.SetBody("text/html", e.Body)

	smtpPort := helper.GetEnv("SMTP_PORT", "")
	intSmptPort, _ := strconv.Atoi(smtpPort)
	dialer := gomail.NewDialer(
		helper.GetEnv("SMTP_HOST", ""),
		intSmptPort,
		helper.GetEnv("SMTP_AUTH_EMAIL", ""),
		helper.GetEnv("SMTP_AUTH_PASSWORD", ""),
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	return nil
}
