package main

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"photo-sharing/db"
	"photo-sharing/handler"
	"photo-sharing/middleware"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	mw := echoview.NewMiddleware(goview.Config{
		Root:         "views",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: true,
	})

	e.Use(mw)
	e.Static("/css", "static/css")
	e.Static("/js", "static/js")
	e.Static("/uploads", "uploads")

	if err := db.Open(); err != nil {
		e.Logger.Fatal("Could not connect to database")
	}

	db.AutoMigrate()

	e.GET("/", handler.GetHomepage, middleware.IsAuthenticated)
	e.GET("/login", handler.GetLogin)
	e.POST("/login", handler.PostLogin)
	e.GET("/register", handler.GetRegister)
	e.POST("/register", handler.PostRegister)
	e.GET("/logout", handler.GetLogout)
	e.POST("/post", handler.PostPost, middleware.IsAuthenticated)
	e.GET("/group/:id", handler.GetGroup, middleware.IsAuthenticated)
	e.POST("/group/:id/invite", handler.PostGroupSendInvite, middleware.IsAuthenticated)
	e.POST("/group/create", handler.PostGroupCreate, middleware.IsAuthenticated)

	e.Logger.Fatal(e.Start(":8080"))
}
