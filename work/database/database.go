package database

import (
	"log"
	"os"
	"github.com/cosmovid21/fiber/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var Database DBInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect \n", err.Error())
		os.Exit(2)
	}

	log.Println("connected to database sucessfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
		db.AutoMigrate(models.User{}, &models.Product{}, &models.Order{})
	Database = DBInstance{Db: db}
}