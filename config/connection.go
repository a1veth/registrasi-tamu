package config

import (
	"fmt"
	"log"
	"os"

	"registrasi-tamu/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	dbname := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"),
	)

	database, err := gorm.Open(postgres.Open(dbname), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	database.AutoMigrate(&models.Tamu{}, &models.Admin{})
	DB = database
	fmt.Println("Database connected and migrated")
}