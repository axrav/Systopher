package helpers

import (
	"fmt"

	"github.com/axrav/Systopher/backend/db"
	"github.com/axrav/Systopher/backend/errors"
	"github.com/axrav/Systopher/backend/models"
)

func GetUserData(email string) *models.UserData {
	var user db.User
	err := db.Pgres.Where("email = ?", email).First(&user).Error
	if err != nil {
		fmt.Println(err)
		return &models.UserData{}
	}

	key := GenerateId("WS-")
	err = SetUserId(key, email)
	if err != nil {
		fmt.Println(err)
		return &models.UserData{}
	}
	return &models.UserData{
		Email:     user.Email,
		Username:  user.Username,
		UniqueID:  user.UniqueID,
		Servers:   GetServers(email),
		RandomKey: key,
	}

}

func GetEmailFromId(token string) string {
	email, err := db.RedisClient.Get(db.Ctx, token).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return email
}

func CheckUserExists(email string) error {
	var user db.User
	err := db.Pgres.Where(&db.User{IsVerified: true, Email: email}).Find(&db.User{}).First(&user).Error
	if err != nil {
		fmt.Println("Error in CheckUserExists")
		fmt.Println(err)
		return err
	}

	if user.Email == email {
		return nil
	} else {
		return errors.InvalidUser.Error()
	}
}

func CreateUser(email string, hash string, username string, u_id string) error {
	err := db.Pgres.Create(&db.User{Email: email, Password: hash, Username: username, UniqueID: u_id, IsVerified: false}).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SetUserId(u_id string, email string) error {
	err := db.RedisClient.Set(db.Ctx, u_id, email, 0).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
