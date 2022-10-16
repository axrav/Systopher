package helpers

import (
	"os"

	gomail "gopkg.in/gomail.v2"
)

func SendMail(email string, subject string, body string) error {
	msg := gomail.NewMessage()
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	from := os.Getenv("GMAIL_ID")
	password := os.Getenv("GMAIL_PASSWORD") // app specific password
	msg.SetHeader("From", from)
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	d := gomail.NewDialer(smtpHost, smtpPort, from, password)
	err := d.DialAndSend(msg)
	if err != nil {
		return err
	}
	return nil

}
