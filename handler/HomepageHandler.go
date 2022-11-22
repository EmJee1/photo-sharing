package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetHomepage(context echo.Context) error {
	return context.Render(http.StatusOK, "Homepage.html", echo.Map{})
}
