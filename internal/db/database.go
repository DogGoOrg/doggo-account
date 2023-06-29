package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres" //nolint
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
