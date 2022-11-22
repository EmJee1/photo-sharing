package main

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"log"
	"photo-sharing/db"
	"photo-sharing/handler"
)

func main() {
	e := echo.New()

	e.Renderer = echoview.Default()

	if err := db.Open(); err != nil {
		log.Fatal("Could not connect to database")
	}

	db.AutoMigrate()

	// TODO: add is-logged-in middleware
	e.GET("/", handler.GetHomepage)
	e.GET("/login", handler.GetLogin)
	e.POST("/login", handler.PostLogin)
	e.GET("/register", handler.GetRegister)
	e.POST("/register", handler.PostRegister)

	e.Logger.Fatal(e.Start(":8080"))
}
