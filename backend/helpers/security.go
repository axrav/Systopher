package helpers

import (
	"os"
	"time"

	"github.com/axrav/Systopher/backend/db"
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
	rows, err := db.Pgres.Query(`SELECT "password" FROM USERS where email=$1`, email)
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

func GenerateJWT(email string, remember bool, forType string) (string, error) {
	var claims *jwt.MapClaims
	if remember {
		claims = &jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 360).Unix(), // 15 days expiration time
		}
	} else {
		claims = &jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 1).Unix(), // 1 hour expiration time
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var secret string
	if forType == "browse" {
		secret = os.Getenv("JWT_SECRET")
	} else {
		secret = os.Getenv("FORGET_SECRET")
		claims = &jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Minute * 10).Unix(), // 10 minutes expiration time
		}
	}
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
