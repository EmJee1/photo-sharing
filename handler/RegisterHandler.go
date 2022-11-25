package handler

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/model"
	"photo-sharing/util"
)

func GetRegister(context echo.Context) error {
	return echoview.Render(context, http.StatusOK, "Register", echo.Map{
		"title": "Login",
	})
}

func PostRegister(context echo.Context) error {
	email := context.FormValue("email")
	password := context.FormValue("password")

	passwordHash, err := util.HashPassword(password)
	if err != nil {
		return echoview.Render(context, http.StatusInternalServerError, "Register", echo.Map{
			"title": "Register",
		})
	}

	db.DB.Create(&model.User{Email: email, Password: passwordHash})

	// TODO: auto-login after account creation
	return echoview.Render(context, http.StatusCreated, "Login", echo.Map{
		"title": "Login",
	})
}
