package helpers

import (
	"fmt"

	"github.com/axrav/Systopher/backend/db"
	"github.com/axrav/Systopher/backend/types"
)

func GetUserData(email string) *types.UserData {
	row, err := db.Pgres.Query("SELECT username,uniqueid FROM users where email=$1", email)
	if err != nil {
		fmt.Println(err)
	}
	var username string
	var uniqueID string
	for row.Next() {
		row.Scan(&username, &uniqueID)
	}
	key := GenerateId("WS-")
	err = SetUserId(key, email)
	if err != nil {
		fmt.Println(err)
		return &types.UserData{}
	}
	return &types.UserData{
		Email:     email,
		Username:  username,
		UniqueID:  uniqueID,
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
	rows, err := db.Pgres.Query(`SELECT "email" FROM users where email=$1`, email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var Email string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&Email)
	}
	if Email == email {
		return nil
	} else {
		return fmt.Errorf("user does not exist")
	}
}
