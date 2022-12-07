package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"photo-sharing/model"
)

var DB *gorm.DB

func getDsn() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, name, port)
}

func Open() error {
	db, err := gorm.Open(postgres.Open(getDsn()), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func AutoMigrate() {
	err := DB.AutoMigrate(&model.User{}, &model.Group{}, &model.GroupInvite{}, &model.Post{})
	if err != nil {
		log.Fatal("Failed to execute auto-migrate on DB")
	}
}
