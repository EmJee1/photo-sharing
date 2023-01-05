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
	return echoview.Render(context, http.StatusOK, "register", echo.Map{
		"title":      "Registreer",
		"hideNavbar": true,
	})
}

func PostRegister(context echo.Context) error {
	email := context.FormValue("email")
	password := context.FormValue("password")
	username := context.FormValue("username")

	passwordHash, err := util.HashPassword(password)
	if err != nil {
		return echoview.Render(context, http.StatusInternalServerError, "register", echo.Map{
			"title":      "Register",
			"hideNavbar": true,
		})
	}

	err = db.DB.Create(&model.User{Email: email, Password: passwordHash, Username: username}).Error
	if err != nil {
		return echoview.Render(context, http.StatusBadRequest, "register", echo.Map{
			"error":      "Controleer de velden en probeer opnieuw",
			"title":      "Register",
			"hideNavbar": true,
		})
	}

	// TODO: auto-login after account creation
	return context.Redirect(http.StatusSeeOther, "/login")
}
