package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		logrus.Error("Error loading .env file", err)
	}

	dsn := os.Getenv("POSTGRES_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("Error connecting to database", err)
	}
	DB = db
	logrus.Info("Connected to database")
}
