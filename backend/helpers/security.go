package helpers

import (
	"os"
	"time"

	"github.com/axrav/SysAnalytics/backend/db"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CompareHashAndPassword(password string, email string) (bool, error) {
	rows, err := db.Db.Query(`SELECT "password" FROM USERS where email=$1`, email)
	if err != nil {
		return false, err
	}
	var hashofPassword string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&hashofPassword)
	}
	check := CheckPasswordHash(password, hashofPassword)
	return check, nil

}

func GenerateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 2).Unix(), // 2 hours expiration time
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
