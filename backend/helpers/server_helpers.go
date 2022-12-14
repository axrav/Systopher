package helpers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/axrav/Systopher/backend/db"
	"github.com/axrav/Systopher/backend/errors"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func SaveServerToken(ip string, token string) bool {
	if err := db.RedisClient.Set(db.Ctx, ip, token, 0).Err(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func GenerateId(typeof string) string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return typeof + string(b)
}

func CheckServerAndDelete(server *db.Server) error {
	var Ipaddr string
	err := db.Pgres.Find(&db.Server{Ip: server.Ip, Owner: server.Owner}).Scan(&Ipaddr).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	// delete server from database
	err = db.Pgres.Delete(&db.Server{Ip: server.Ip, Owner: server.Owner}).Error
	if err != nil {
		fmt.Println(err)
		return err

	}
	return nil
}

func GetServers(email string) []db.Server {
	var servers []db.Server
	err := db.Pgres.Find(&db.Server{Owner: db.User{Email: email}}).Scan(&servers).Error
	if err != nil {
		fmt.Println(err)
	}
	return servers
}

func CheckServerAndAdd(server *db.Server) error {
	var Ipaddr string
	err := db.Pgres.Find(&db.Server{Ip: server.Ip}).Error
	if err != nil {
		fmt.Println(err)
		return err
	}

	if Ipaddr == server.Ip {
		return errors.AlreadyExists.Error()
	} else {
		test, err := TestRequest(server.Ip, server.Port, server.Token) // perform a test request
		if test {
			// insert server to database
			err = db.Pgres.Create(&db.Server{Ip: server.Ip, Port: server.Port, Token: server.Token, Owner: server.Owner}).Error
			if err != nil {
				fmt.Println(err)
				return err
			}
			out := SaveServerToken("http://"+server.Ip+":"+server.Port, server.Token)
			if out {
				return nil
			} else {
				return errors.InternalServerError.Error()
			}
		} else {
			return err
		}
	}

}
