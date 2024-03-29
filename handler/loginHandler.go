package handler

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/model"
	"photo-sharing/repository"
	"photo-sharing/util"
	"strconv"
)

func GetLogin(context echo.Context) error {
	return echoview.Render(context, http.StatusOK, "login", echo.Map{
		"title":      "Login",
		"hideNavbar": true,
	})
}

func PostLogin(context echo.Context) error {
	email := context.FormValue("email")
	password := context.FormValue("password")

	user := &model.User{}
	err := repository.GetUserByEmail(email, &user)
	if err != nil || !util.CheckPasswordHash(password, user.Password) {
		return echoview.Render(context, http.StatusUnauthorized, "login", echo.Map{
			"error":      "Invalid username & password combination",
			"title":      "Login",
			"hideNavbar": true,
		})
	}

	expiresAt, token, err := util.GenerateJWT(strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		return echoview.Render(context, http.StatusInternalServerError, "login", echo.Map{
			"error":      "Something unexpected went wrong",
			"title":      "Login",
			"hideNavbar": true,
		})
	}

	context.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  expiresAt,
		HttpOnly: true,
	})

	context.Redirect(http.StatusFound, "/")
	return nil
}
