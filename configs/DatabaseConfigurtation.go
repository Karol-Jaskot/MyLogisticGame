package configs

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var FileName = "database/MainDB.db"

func GetConnection() *gorm.DB {
	if _, err := os.Stat(FileName); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll("./database", 0755)
		os.Create(FileName)
	} else {
		// file exists
	}

	db, err := gorm.Open(sqlite.Open(FileName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func SetTstConnection() {
	FileName = "database/MainDB_tst.db"
}
