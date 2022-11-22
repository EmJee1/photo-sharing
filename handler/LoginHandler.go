package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetLogin(context echo.Context) error {
	return context.Render(http.StatusOK, "Login.html", echo.Map{})
}

func PostLogin(context echo.Context) error {
	return context.Render(http.StatusOK, "Login.html", echo.Map{})
}
