package app

import (
	"errors"
	"os"

	"github.com/emcassi/open-stash-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() error {
	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		return errors.New("environment variable 'DB_FILE' required and not found")
	}

	var err error 
	Db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		return err
	}

	Db.AutoMigrate(&models.User{})

	return nil
}
