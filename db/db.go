package db

import (
	"os"

	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDB() *gorm.DB {
	url := os.Getenv("DATABASE_URL") + os.Getenv("SSL_MODE")
	db, err := gorm.Open("postgres", url)

	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	return db
}
