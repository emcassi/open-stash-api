package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/emcassi/open-stash-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() error {

	host := os.Getenv("DB_HOST")
	if host == "" {
		return errors.New("required environment variable 'DB_HOST' not found")
	}

	username := os.Getenv("DB_USER")
	if host == "" {
		return errors.New("required environment variable 'DB_USER' not found")
	}

	password := os.Getenv("DB_PASSWORD")
	if host == "" {
		return errors.New("required environment variable 'DB_PASSWORD' not found")
	}

	dbName := os.Getenv("DB_NAME")
	if host == "" {
		return errors.New("required environment variable 'DB_NAME' not found")
	}

	port := os.Getenv("DB_PORT")
	if host == "" {
		return errors.New("required environment variable 'DB_PORT' not found")
	}

	ssl := os.Getenv("DB_SSL")
	if host == "" {
		return errors.New("required environment variable 'DB_SSL' not found")
	}

	timezone := os.Getenv("DB_TIMEZONE")
	if host == "" {
		return errors.New("required environment variable 'DB_TIMEZONE' not found")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, username, password, dbName, port, ssl, timezone)

	var err error 
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	Db.AutoMigrate(&models.User{})

	return nil
}
