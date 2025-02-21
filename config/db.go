package config

import (
	"log"
	"os"

	"food-delivery/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected!")

	// Auto-migrate models
	if err = DB.AutoMigrate(&models.User{}, &models.Restaurant{}, &models.Order{}); err != nil {
		log.Fatal("Error during migration:", err)
	}

	log.Println("Database migrated successfully!")
}
