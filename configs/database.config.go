package configs

import (
	"OneTix/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	env := LoadEnv()

	config := env.DBURL

	log.Println("Connecting to database...")

	database, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")

	err = database.AutoMigrate(&models.MstUser{}, &models.MstOrganizer{}, &models.MstEvent{}, &models.TrsTicket{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = database
}
