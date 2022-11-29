package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetLogout(context echo.Context) error {
	context.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})

	return context.Redirect(http.StatusSeeOther, "/login")
}
