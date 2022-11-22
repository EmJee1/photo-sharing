package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	// TODO: add is-logged-in middleware
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Slash!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
