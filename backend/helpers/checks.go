package helpers

import (
	"fmt"
	"strings"
	"unicode"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/axrav/Systopher/backend/db"
	"github.com/axrav/Systopher/backend/errors"
	"github.com/axrav/Systopher/backend/types"
)

var (
	verifier = emailverifier.NewVerifier()
)

func UserCheckers(user *types.User) error {
	res, err := verifier.Verify(user.Email)
	if err != nil {
		return err
	}
	if !res.Syntax.Valid {
		return errors.InvalidEmail.Error()
	}
	userNameExists := CheckUserNameExists(user.Username)
	if userNameExists != nil {
		return userNameExists
	}
	passwordValidator := CheckPassword(user.Password)
	if !passwordValidator {
		return errors.InvalidPassword.Error()
	}
	return nil

}

func CheckUserNameExists(username string) error {

	if len(strings.Split(username, " ")) > 1 {
		return errors.InvalidUsername.Error()
	}
	rows, err := db.Pgres.Query(`SELECT "username" FROM users where username=$1`, username)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var Username string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&Username)
	}
	if Username == username {
		return errors.UsernameTaken.Error()
	} else {
		return nil
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
