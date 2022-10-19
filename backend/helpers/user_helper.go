package helpers

import (
	"fmt"

	"github.com/axrav/Systopher/backend/db"
	"github.com/axrav/Systopher/backend/types"
)

func GetUserData(email string) *types.UserData {
	row, err := db.Db.Query("SELECT username,uniqueid FROM users where email=$1", email)
	if err != nil {
		fmt.Println(err)
	}
	var username string
	var uniqueID string
	for row.Next() {
		row.Scan(&username, &uniqueID)
	}

	return &types.UserData{
		Email:    email,
		Username: username,
		UniqueID: uniqueID,
		Servers:  GetServers(email),
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
