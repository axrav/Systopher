package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

// postgres database
var Db *sql.DB

func Init() {
	var err error
	var DbURI = os.Getenv("DB_URI")
	Db, err = sql.Open("postgres", DbURI)
	if err != nil {
		panic(err)
	}
	// create tables
	// createUsersTable
	Db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, email TEXT NOT NULL, password TEXT NOT NULL, servers TEXT)")
	// createServersTable
	Db.Exec("CREATE TABLE IF NOT EXISTS servers (id SERIAL PRIMARY KEY, name TEXT NOT NULL, ip TEXT NOT NULL, port TEXT NOT NULL, owner TEXT NOT NULL)")

}
