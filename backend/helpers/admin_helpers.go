package helpers

import (
	"fmt"

	"github.com/axrav/Systopher/backend/db"
	"github.com/axrav/Systopher/backend/models"
)

func GetAdmin(email string) bool {
	row, err := db.Pgres.Query(`SELECT "email" FROM admins where email=$1`, email)
	if err != nil {
		fmt.Println(err)
		return false
	}
	var Email string
	defer row.Close()
	for row.Next() {
		row.Scan(&Email)
	}
	if Email == email {
		return true
	} else {
		return false
	}
}

func GetAllUsers() ([]models.UserData, error) {
	var users []models.UserData
	rows, err := db.Pgres.Query(`SELECT "email" FROM users`)
	for rows.Next() {
		var email string
		rows.Scan(&email)
		users = append(users, *GetUserData(email))
	}

	if err != nil {
		fmt.Println(err)
		return users, err
	}
	return users, nil
}

func AddAdmin(email string) error {
	_, err := db.Pgres.Exec(`INSERT INTO admins(email) VALUES($1)`, email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
