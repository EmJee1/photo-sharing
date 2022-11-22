package main

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"photo-sharing/handlers"
	"time"
)

type User struct {
	ID        uint
	UserName  string
	CreatedAt time.Time
}

func main() {
	e := echo.New()

	e.Renderer = echoview.Default()

	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	db.AutoMigrate(&User{})
	db.Create(&User{UserName: "Mart-Jan"})

	// TODO: add is-logged-in middleware
	e.GET("/", handlers.Homepage)

	e.Logger.Fatal(e.Start(":8080"))
}
