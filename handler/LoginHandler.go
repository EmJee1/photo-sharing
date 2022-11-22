package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetLogin(context echo.Context) error {
	return context.Render(http.StatusOK, "Login.html", echo.Map{})
}

func PostLogin(context echo.Context) error {
	email := context.FormValue("email")
	password := context.FormValue("password")

	fmt.Printf("Email: %s \n", email)
	fmt.Printf("Password: %s \n", password)

	return context.Render(http.StatusOK, "Login.html", echo.Map{})
}
