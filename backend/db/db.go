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

	// schema
	// create tables
	// createUsersTable
	_, err = Db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, email TEXT NOT NULL, password TEXT NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, Unique(email))")
	// createServersTable
	if err != nil {
		panic(err)
	}
	_, err = Db.Exec("CREATE TABLE IF NOT EXISTS servers (id SERIAL PRIMARY KEY, name TEXT, ip TEXT NOT NULL UNIQUE, port TEXT NOT NULL, owner TEXT NOT NULL REFERENCES users (email), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		panic(err)
	}

}
