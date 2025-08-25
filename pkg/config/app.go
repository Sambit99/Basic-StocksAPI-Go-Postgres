package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	POSTGRES_URL := os.Getenv("POSTGRES_URL")

	database, err := gorm.Open(postgres.Open(POSTGRES_URL), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting to Database", err.Error())
	}

	fmt.Println("Connected to Database ✅")

	db = database
}

func GetDB() *gorm.DB {
	return db
}
