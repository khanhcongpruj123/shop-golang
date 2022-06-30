package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
