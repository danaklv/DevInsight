package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	settings := "host=localhost user=postgres password=dana1234 dbname=DevInsight port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(settings))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDatabase() {
	var err error
	DB, err = ConnectDatabase()
	if err != nil {
		log.Fatal("Error conect to databse: ", err)

	}
}
