// package helpers
package helpers

import (
	"github.com/axrav/Systopher/backend/db"
	"github.com/axrav/Systopher/backend/errors"
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
		return errors.InvalidCred.Error()
	}

}

func UpdatePassword(password string, email string) error {
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}
	err = db.Pgres.Model(&db.User{}).Where("email = ?", email).Update("password", hash).Error
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
		return errors.InternalServerError.Error()
	}

	return nil

}
