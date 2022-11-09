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

// func CheckUserExists(email string) bool {
// 	row, err := db.Db.Query("SELECT email FROM users where email=$1", email)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for row.Next() {
// 		return true
// 	}
// 	return false
// }

// func ForgetPasswordAndSendOTP(email string) bool {
// 	if CheckUserExists(email) {
// 		otp, err := GenerateOTP()
// 		if err != nil {
// 			fmt.Println(err)
// 			return false
// 		}
// 		err = db.RedisClient.Set(db.Ctx, email, otp, 0).Err()
// 		if err != nil {
// 			fmt.Println(err)
// 			return false
// 		}
// 		sent := SendForgetMail(email, otp)
// 		if sent {
// 			return true
// 		}

// 	}
// 	return false
// }
