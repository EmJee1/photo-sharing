package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"photo-sharing/model"
)

var DB *gorm.DB

func Open() error {
	// TODO: extract database credentials in .env file
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func AutoMigrate() {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to execute auto-migrate on DB")
	}
}
