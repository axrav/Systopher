package helpers

import (
	"fmt"

	"github.com/axrav/SysAnalytics/backend/db"
	"github.com/axrav/SysAnalytics/backend/types"
)

func GetUserData(email string) *types.UserData {
	row, err := db.Db.Query("SELECT username FROM users where email=$1", email)
	if err != nil {
		fmt.Println(err)
	}
	var username string
	row.Scan(&username)

	return &types.UserData{
		Email:    email,
		Username: username,
		Servers:  GetServers(email),
	}

}

// Get user data from database
