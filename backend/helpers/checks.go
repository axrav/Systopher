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
	rows, err := db.Pgres.Query(`SELECT "username" FROM users where username=$1`, username)
	if err != nil {
		fmt.Println(err)
		return errors.InternalServerError.Merror()
	}
	var Username string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&Username)
	}
	if Username == username {
		return errors.UsernameTaken.Merror()
	} else {
		return errors.Error{}
	}

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
