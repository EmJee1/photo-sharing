package main

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"photo-sharing/db"
	"photo-sharing/handler"
	"photo-sharing/middleware"
	"photo-sharing/util"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	mw := echoview.NewMiddleware(util.GoviewConfig)

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
	e.POST("/group", handler.PostGroup, middleware.IsAuthenticated)
	e.GET("/group/:id", handler.GetGroup, middleware.IsAuthenticated, middleware.IsGroupUser)
	e.GET("/settings", handler.GetSettings, middleware.IsAuthenticated)

	// Routes in the API group should always return a JSON response
	api := e.Group("/api")

	api.POST("/post", handler.PostPost, middleware.IsAuthenticated)
	api.DELETE("/post", handler.DeletePost, middleware.IsAuthenticated)
	api.GET("/invite", handler.GetInvites, middleware.IsAuthenticated)
	api.POST("/invite", handler.PostInvite, middleware.IsAuthenticated, middleware.IsGroupUser)
	api.POST("/invite/respond", handler.PostInviteRespond, middleware.IsAuthenticated)
	api.POST("/like", handler.PostLike, middleware.IsAuthenticated)
	api.POST("/comment", handler.PostComment, middleware.IsAuthenticated)
	api.DELETE("/comment", handler.DeleteComment, middleware.IsAuthenticated)

	e.Logger.Fatal(e.Start(":8080"))
}
