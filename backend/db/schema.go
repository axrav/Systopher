package db

import "gorm.io/gorm"

type Server struct {
	gorm.Model
	Ip    string `json:"ip_address"`
	Port  string `json:"port"`
	Name  string `json:"name" `
	Owner User   `gorm:"foreignKey:Email;references:Ip"`
	Token string `json:"token" `
}

type User struct {
	gorm.Model
	Username   string `gorm:"unique"`
	Email      string `gorm:"unique"`
	Password   string
	UniqueID   string
	IsVerified bool
	Servers    []Server `gorm:"foreignKey:Ip;references:Username"`
}

type Admin struct {
	gorm.Model
	User User `gorm:"foreignKey:Email;references:ID"`
}
