package helpers

import (
	"fmt"

	"github.com/axrav/SysAnalytics/backend/db"
	"github.com/axrav/SysAnalytics/backend/types"
)

func CheckServerAndAdd(server *types.Server) error {
	rows, err := db.Db.Query(`SELECT "ip" FROM servers where ip=$1 and owner=$2`, server.Ip, server.Owner)
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
		return fmt.Errorf("server already exists")
	} else {
		test, err := TestRequest(server.Ip, server.Port, server.Token) // perform a test request
		if test {
			// insert server to database
			_, err = db.Db.Exec(`INSERT INTO servers (name, ip, port, owner, token) VALUES ($1, $2, $3, $4, $5)`, server.Name, server.Ip, server.Port, server.Owner, server.Token)
			if err != nil {
				fmt.Println(err)
				return err
			}
			out := SaveServerToken("http://"+server.Ip+":"+server.Port, server.Token)
			if out {
				return nil
			} else {
				return fmt.Errorf("error saving server token")
			}
		} else {
			return err
		}
	}

}

func CheckServerAndDelete(server *types.Server) error {

	rows, err := db.Db.Query(`SELECT "ip" FROM servers where ip=$1`, server.Ip)
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
		_, err = db.Db.Exec(`DELETE FROM servers where ip=$1 and owner=$2`, server.Ip, server.Owner)
		if err != nil {
			fmt.Println(err)
			return err
		}

	}
	return nil
}

func GetServers(email string) []string {
	rows, err := db.Db.Query(`SELECT ip,port FROM servers where owner=$1`, email)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var servers []string
	for rows.Next() {
		var ip string
		var port string
		rows.Scan(&ip, &port)
		servers = append(servers, "http://"+ip+":"+port)
	}
	return servers
}
