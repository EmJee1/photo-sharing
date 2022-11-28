package handler

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/model"
)

func GetHomepage(context echo.Context) error {
	return echoview.Render(context, http.StatusOK, "homepage", echo.Map{
		"title": "Overview",
		"user":  model.User{Email: "john@doe.com"},
	})
}
