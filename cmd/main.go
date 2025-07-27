package main

import (
	"athenify/domain"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Println("Loading environment variables")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file found")
	}
	postgresUsername := os.Getenv("POSTGRES_USERNAME")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	postgresIP := os.Getenv("POSTGRES_IP")
	postgresPort := os.Getenv("POSTGRES_PORT")

	log.Println("Connecting to database")
	source := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", postgresUsername, postgresPassword, postgresIP, postgresPort, postgresDB)
	db, err := gorm.Open(postgres.Open(source), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Migrating")
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Task{})
}
