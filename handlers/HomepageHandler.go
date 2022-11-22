package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Homepage(context echo.Context) error {
	return context.Render(http.StatusOK, "Homepage.html", echo.Map{"name": "Homepage :-)"})
}
