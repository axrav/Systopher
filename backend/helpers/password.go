// package helpers
package helpers

import (
	"fmt"

	"github.com/axrav/Systopher/backend/db"
)

func ChangePassword(email string, oldPassword string, newPassword string) error {
	compare, err := CompareHashAndPassword(oldPassword, email)
	if err != nil {
		return err
	}
	if compare {
		err := UpdatePassword(newPassword, email)
		if err != nil {
			return err
		}
		return nil
	} else {
		return fmt.Errorf("wrong password")
	}

}

func UpdatePassword(password string, email string) error {
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}

	_, err = db.Pgres.Exec("UPDATE users SET password = $1 WHERE email = $2", hash, email)
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

	go SendMail(email, "Password Reset", "Hello, <br> Your OTP for resetting Systopher's password is: <b>"+otp+"</b>")
	success := SaveOtp(email, otp)
	if !success {
		return fmt.Errorf("error while saving otp")
	}

	return nil

}
