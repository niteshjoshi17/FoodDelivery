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
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Warning: .env file not found, using system environment variables")
	}

	// Get the DATABASE_URL from the environment
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Error: DATABASE_URL is not set in environment variables")
	}

	// Connect to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully!")

	// Auto-migrate models (Removed models.User{})
	if err = DB.AutoMigrate(&models.Restaurant{}, &models.Order{}); err != nil {
		log.Fatal("Error during migration:", err)
	}

	log.Println("Database migrated successfully!")
}
