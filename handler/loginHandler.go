package handler

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/model"
	"photo-sharing/util"
)

func GetLogin(context echo.Context) error {
	return echoview.Render(context, http.StatusOK, "login", echo.Map{
		"title": "Login",
	})
}

func PostLogin(context echo.Context) error {
	email := context.FormValue("email")
	password := context.FormValue("password")

	user := &model.User{}
	db.DB.Where("email = ?", email).First(&user)
	if user == nil || !util.CheckPasswordHash(password, user.Password) {
		return echoview.Render(context, http.StatusUnauthorized, "login", echo.Map{
			"error": "Invalid username & password combination",
			"title": "Login",
		})
	}

	// TODO: supply user with token and redirect to homepage
	return echoview.Render(context, http.StatusOK, "login", echo.Map{
		"title": "Login",
	})
}
