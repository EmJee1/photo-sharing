package handler

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/model"
)

func GetHomepage(context echo.Context) error {
	user := &model.User{}
	db.DB.Where("id = ?", context.Get("userId")).First(&user)
	return echoview.Render(context, http.StatusOK, "homepage", echo.Map{
		"title": "Overview",
		"user":  user,
	})
}
