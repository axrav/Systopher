package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

// postgres database
var Pgres *sql.DB

func InitPostgres() {
	var err error
	var DbURI = os.Getenv("POSTGRES_DB_URI")
	Pgres, err = sql.Open("postgres", DbURI)
	if err != nil {
		panic(err)
	}

	// schema
	// create tables
	// createUsersTable
	_, err = Pgres.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY,username TEXT UNIQUE, email TEXT NOT NULL UNIQUE,UniqueID TEXT NOT NULL UNIQUE, password TEXT NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, isverified BOOLEAN DEFAULT FALSE)")
	// createServersTable
	if err != nil {
		panic(err)
	}
	_, err = Pgres.Exec("CREATE TABLE IF NOT EXISTS servers (id SERIAL PRIMARY KEY, name TEXT, ip TEXT NOT NULL , port TEXT NOT NULL, owner TEXT NOT NULL REFERENCES users (email),token TEXT NOT NULL UNIQUE, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		panic(err)
	}

}
