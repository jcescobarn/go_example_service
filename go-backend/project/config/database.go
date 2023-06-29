package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnvFile() {
	enverror := godotenv.Load(".env")
	if enverror != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectToDB() {
	var err error

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSL_MODE")

	database_string := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, db_name, port, sslmode)
	DB, err = gorm.Open(postgres.Open(database_string), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB")
	}
}
