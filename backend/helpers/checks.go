package helpers

import (
	"fmt"
	"regexp"
	"strings"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/axrav/Systopher/backend/db"
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
		return fmt.Errorf("invalid email")
	}
	userNameExists := CheckUserNameExists(user.Username)
	if userNameExists != nil {
		return userNameExists
	}
	passwordValidator := CheckPassword(user.Password)
	if passwordValidator != nil {
		return fmt.Errorf("password is not valid")
	}
	return nil

}

func CheckUserNameExists(username string) error {

	if len(strings.Split(username, " ")) > 1 {
		return fmt.Errorf("username should not contain spaces")
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
		return fmt.Errorf("username already exists")
	} else {
		return nil
	}

}

func CheckPassword(password string) error {
	//restr := regexp.QuoteMeta()
	// fmt.Println(restr)
	re, _ := regexp.Compile(`\A(?:[A-Z]|[a-z]|[0-9]|.){8,}`)
	fmt.Println(password)
	if re.MatchString(password) {
		return nil
	} else {
		return fmt.Errorf("password is not valid")
	}
}
