package helpers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/axrav/Systopher/backend/db"
	"github.com/axrav/Systopher/backend/errors"
	"github.com/axrav/Systopher/backend/types"
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

func CheckServerAndDelete(server *types.Server) error {

	rows, err := db.Pgres.Query(`SELECT "ip" FROM servers where ip=$1`, server.Ip)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var Ipaddr string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&Ipaddr)
	}
	if Ipaddr == server.Ip {
		// delete server from database
		_, err = db.Pgres.Exec(`DELETE FROM servers where ip=$1 and owner=$2`, server.Ip, server.Owner)
		if err != nil {
			fmt.Println(err)
			return err
		}

	}
	return nil
}

func GetServers(email string) []types.Server {
	rows, err := db.Pgres.Query(`SELECT ip,port,token,name,owner FROM servers where owner=$1`, email)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var servers []types.Server
	for rows.Next() {
		var server types.Server
		err = rows.Scan(&server.Ip, &server.Port, &server.Token, &server.Name, &server.Owner)
		if err != nil {
			fmt.Println(err)
		}
		servers = append(servers, server)
	}
	return servers
}

func CheckServerAndAdd(server *types.Server) error {
	rows, err := db.Pgres.Query(`SELECT "ip" FROM servers where ip=$1 and owner=$2`, server.Ip, server.Owner)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var Ipaddr string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&Ipaddr)
	}
	if Ipaddr == server.Ip {
		return errors.AlreadyExists.Error()
	} else {
		test, err := TestRequest(server.Ip, server.Port, server.Token) // perform a test request
		if test {
			// insert server to database
			_, err = db.Pgres.Exec(`INSERT INTO servers (name, ip, port, owner, token) VALUES ($1, $2, $3, $4, $5)`, server.Name, server.Ip, server.Port, server.Owner, server.Token)
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
