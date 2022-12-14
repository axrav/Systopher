package helpers

import (
	"fmt"
	"strings"
	"unicode"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/axrav/Systopher/backend/db"
	"github.com/axrav/Systopher/backend/errors"
	"github.com/axrav/Systopher/backend/models"
)

var (
	verifier = emailverifier.NewVerifier()
)

func UserCheckers(user *models.User) errors.Error {
	res, err := verifier.Verify(user.Email)
	if err != nil {
		return errors.InvalidEmail.Merror()
	}
	if !res.Syntax.Valid {
		return errors.InvalidEmail.Merror()
	}
	userNameExists := CheckUserNameExists(user.Username)
	if userNameExists.Err != nil {
		return userNameExists
	}
	passwordValidator := CheckPassword(user.Password)
	if !passwordValidator {
		return errors.InvalidPassword.Merror()
	}
	return errors.Error{}

}

func CheckUserNameExists(username string) errors.Error {

	if len(strings.Split(username, " ")) > 1 {
		return errors.InvalidUsername.Merror()
	}
	var user db.User
	err := db.Pgres.Where("username = ?", username).First(&user).Error
	if err != nil {
		fmt.Println(err)
		return errors.Error{}
	}
	return errors.UsernameTaken.Merror()

}

func CheckPassword(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(password) >= 7 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
