package handler

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/model"
	"photo-sharing/repository"
)

func GetHomepage(context echo.Context) error {
	user := &model.User{}
	repository.GetUser(
		context.Get("userId").(uint),
		&user,
		"Groups.Posts",
	)

	return echoview.Render(context, http.StatusOK, "homepage", echo.Map{
		"title": "Overview",
		"user":  user,
	})
}
