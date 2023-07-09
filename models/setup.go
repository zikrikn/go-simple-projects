package models

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

// ConnectDB connects to the database
func ConnectDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	user_db := os.Getenv("POSTGRES_USER")
	password_db := os.Getenv("POSTGRES_PASSWORD")
	host_db := os.Getenv("POSTGRES_HOST")
	dbname_db := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s  sslmode=verify-full", host_db, user_db, password_db, dbname_db)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})
	// What is AutoMigrate?
	// AutoMigrate digunakan agar bisa ORM-nya digunakan, karena langsung mengintegrasikan dengan database

	DB = db // What is it? DB is a variable
	// Please explain more details?
	// DB is a variable that is used to connect to the database
}
