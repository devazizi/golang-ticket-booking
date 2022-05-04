package services

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}

func ConnectDatabase() Connection {
	var databaseDsn string = os.Getenv("DB_DSN")

	database, dbConnetionErr := gorm.Open(mysql.Open(databaseDsn), &gorm.Config{})

	if dbConnetionErr != nil {
		panic("fail to connect database")
	}

	return Connection{db: database}
}
