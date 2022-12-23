package handler

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/model"
)

func GetGroup(context echo.Context) error {
	groupId := context.Param("id")

	group := &model.Group{}
	err := db.DB.
		Model(&model.Group{}).
		Where("id = ?", groupId).
		Preload("Users").
		Preload("GroupInvites").
		Preload("Posts.User").
		First(&group).
		Error

	if err != nil {
		return echoview.Render(context, http.StatusNotFound, "404", echo.Map{
			"title": "Not Found",
		})
	}

	return echoview.Render(context, http.StatusOK, "group", echo.Map{
		"title": group.Name,
		"group": group,
	})
}
