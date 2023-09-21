package config

import (
	"github.com/joho/godotenv"
	"go-example-api/app/domain/model"
	"gorm.io/gorm/logger"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error when loading .env file. Error: ", err)
	}

	urlConnection := os.Getenv("DB_URL_CONNECTION")

	db, err := gorm.Open(postgres.Open(urlConnection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Error connection database, Error:", err)
	}

	if err = db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("Error when auto migrating, Error:", err)
	}

	return db
}
