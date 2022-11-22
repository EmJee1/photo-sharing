package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Homepage(context echo.Context) error {
	return context.String(http.StatusOK, "Hello, World!")
}
