package handler

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetSettings(context echo.Context) error {
	return echoview.Render(context, http.StatusOK, "settings", echo.Map{})
}
