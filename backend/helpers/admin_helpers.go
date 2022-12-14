package helpers

import (
	"fmt"

	"github.com/axrav/Systopher/backend/db"
)

func GetAdmin(email string) bool {
	var admin db.Admin
	err := db.Pgres.Where("email = ?", email).First(&admin).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func GetAllUsers() ([]db.User, error) {
	var users []db.User
	err := db.Pgres.Find(&users).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil

}

func AddAdmin(email string) error {
	err := db.Pgres.Create(&db.Admin{User: db.User{Email: email}}).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteAdmin(email string) error {
	err := db.Pgres.Delete(&db.Admin{}, "email = ?", email).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
