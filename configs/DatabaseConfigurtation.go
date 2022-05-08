package configs

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func GetConnection() *gorm.DB {
	filename := "./database/MainDB.db"

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll("./database", 0755)
		os.Create(filename)
	} else {
		// file exists
	}

	db, err := gorm.Open(sqlite.Open("database/MainDB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
