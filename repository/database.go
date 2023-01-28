package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"photo-sharing/model"
)

var _connection *gorm.DB

func getDsn() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, name, port)
}

func connection() *gorm.DB {
	if _connection != nil {
		return _connection
	}

	db, err := gorm.Open(postgres.Open(getDsn()), &gorm.Config{})
	if err != nil {
		panic("Database connection failed: " + err.Error())
	}

	if err := db.AutoMigrate(&model.User{}, &model.Group{}, &model.GroupUser{}, &model.Invite{}, &model.Post{}, &model.Comment{}); err != nil {
		panic("Database automigrate failed: " + err.Error())
	}

	_connection = db

	return db
}
