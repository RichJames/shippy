package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func CreateConnection() (*gorm.DB, error) {

	// Get database details from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

//	return gorm.Open(
//		"postgres",
//		fmt.Sprintf(
//			"host=%s user=%s dbname=%s sslmode=disable password=%s",
//			"postgres://%s:%s@%s/%s?sslmode=disable",
//			host, user, DBName, password,
//			user, password, host, DBName
//		),
//	)

	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=disable",
			user, password, host, DBName,
		),
	)
}
