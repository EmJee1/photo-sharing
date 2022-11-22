package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint
	Email     string
	Password  string
	CreatedAt time.Time
}

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
	DB.AutoMigrate(&User{})
}
