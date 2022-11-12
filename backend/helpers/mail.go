package helpers

import (
	"fmt"
	"os"

	gomail "gopkg.in/gomail.v2"
)

func SendMail(email string, subject string, body string) error {
	msg := gomail.NewMessage()
	smtpHost := os.Getenv("SMTP_HOST")      // your own smtp host
	smtpPort := 587                         // change this to your own smtp port, most likely 587 or 465 for ssl
	from := os.Getenv("EMAIL_ID")           // your own email id
	password := os.Getenv("EMAIL_PASSWORD") // app specific password

	// Set E-Mail sender and recipient here
	msg.SetAddressHeader("From", from, "Team Systopher")
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

func SendForgetPasswordEmail(email string) error {
	otp, err := GenerateOTP()
	if err != nil {
		return err
	}

	go SendMail(email, "Password Reset", "Hello, <br> Your OTP for resetting password is: <b>"+otp+"</b>")
	success := SaveOtp(email, otp)
	if !success {
		return fmt.Errorf("error while saving otp")
	}

	return nil

}
