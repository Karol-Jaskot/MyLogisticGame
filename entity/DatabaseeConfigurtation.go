package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database/MainDB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
