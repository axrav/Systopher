package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// postgres database
var Pgres *gorm.DB

func InitPostgres() {
	var err error
	var DbURI = os.Getenv("POSTGRES_DB_URI")
	Pgres, err = gorm.Open(postgres.Open(DbURI), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	Pgres.AutoMigrate(&User{})
	Pgres.AutoMigrate(&Server{})
	Pgres.AutoMigrate(&Admin{})

}
