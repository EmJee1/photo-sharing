package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
)

func GetRegister(context echo.Context) error {
	return context.Render(http.StatusOK, "Register.html", echo.Map{})
}

func PostRegister(context echo.Context) error {
	email := context.FormValue("email")
	password := context.FormValue("password")

	db.DB.Create(&User{Email: email, Password: password})

	return context.Render(http.StatusOK, "Register.html", echo.Map{})
}
