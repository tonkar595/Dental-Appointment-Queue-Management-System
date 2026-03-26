package db

import (
	models "backend/internal/model"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnected() *gorm.DB {
	godotenv.Load()

	dbhost := os.Getenv("DB_HOST")
	dbuser := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbport := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		dbhost,
		dbuser,
		dbpassword,
		dbname,
		dbport,
	)

	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = dbConnection

	fmt.Println("Database connected successfully")

	return DB
}

func AutoMigrate(database *gorm.DB) {
	database.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Patient{},
		&models.Staff{},
		&models.Appointment{},
		&models.Notification{},
	)
}
