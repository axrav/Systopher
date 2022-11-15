package helpers

import (
	"fmt"

	"github.com/axrav/Systopher/backend/db"
)

func CreateUser(email string, hash string, username string, u_id string) error {
	_, err := db.Pgres.Exec("INSERT INTO users (email, password, username, isverified, uniqueid) VALUES ($1, $2, $3, $4, $5)", email, hash, username, false, u_id)
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
