package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/axrav/SysAnalytics/backend/db"
)

const otpChars = "1234567890"

func GenerateOTP() (string, error) {
	buffer := make([]byte, 6)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < 6; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

func SendOtp(email string, otp string) bool {
	var err error
	subject, err := base64.StdEncoding.DecodeString("T1RQIC0gU3lzQW5hbHl0aWNz")
	if err != nil {
		fmt.Println(err)
		return false
	}
	body, err := base64.StdEncoding.DecodeString("T25lIFRpbWUgUGFzc3dvcmQoT1RQKSBmb3IgcmVnaXN0ZXJpbmcgb24gU3lzQW5hbHl0aWNzIGlzOg==")
	if err != nil {
		fmt.Println(err)
		return false
	}
	final_message := string(body) + fmt.Sprintf("<b> %s </b>", otp)
	go SendMail(email, string(subject), final_message)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func SendOtpAndSave(email string) bool {
	otp, err := GenerateOTP()

	if err != nil {
		fmt.Println(err)
		return false
	}
	sent := SendOtp(email, otp)
	if sent {
		hash, err := HashPassword(otp)
		if err != nil {
			fmt.Println("Error in hashing OTP" + err.Error())
			return false

		}
		err = db.RedisClient.Set(db.Ctx, email, hash, 0).Err()
		if err != nil {
			fmt.Println("Error in saving OTP" + err.Error())
			return false
		}
		return true
	} else {
		return false
	}

}

func VerifyOtp(email, otp string) bool {
	hash, err := db.RedisClient.Get(db.Ctx, email).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	check := CheckPasswordHash(otp, hash)
	go SendMail(email, "Welcome to SysAnalytics", "<b>You have successfully signed up to SysAnalytics and your account is verified now.</b> \n\nYou can now login to your account and start using SysAnalytics. \n\n\nThank you for choosing SysAnalytics. Have a great day!") // this message needs to be changed to a better one with proper formatting
	return check
}

func GetVerified(email string) bool {
	rows, err := db.Db.Query(`SELECT isverified FROM users where email=$1`, email)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var verified bool
	for rows.Next() {
		rows.Scan(&verified)
	}
	return verified
}
