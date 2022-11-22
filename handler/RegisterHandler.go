package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/util"
)

func GetRegister(context echo.Context) error {
	return context.Render(http.StatusOK, "Register.html", echo.Map{})
}

func PostRegister(context echo.Context) error {
	email := context.FormValue("email")
	password := context.FormValue("password")

	passwordHash, err := util.HashPassword(password)
	if err != nil {
		return context.Render(http.StatusInternalServerError, "Register.html", echo.Map{})
	}

	db.DB.Create(&User{Email: email, Password: passwordHash})

	// TODO: auto-login after account creation
	return context.Render(http.StatusOK, "Login.html", echo.Map{})
}
