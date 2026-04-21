package external

// import "gopkg.in/gomail.v2"

type External struct{}

func (*External) SendEmail() error {
	// mailer := gomail.NewMessage()
	// mailer.SetHeader("To")
	// mailer.SetHeader("From")
	// mailer.SetHeader("Subject")
	// mailer.SetBody("text/html", "Test Body")
	return nil
}
