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
	db.DB.Model(&model.Group{}).Where("id = ?", groupId).Preload("Users").Preload("GroupInvites").First(&group)
	if group == nil {
		// TODO: 404 page
		context.Redirect(http.StatusSeeOther, "/")
		return nil
	}

	return echoview.Render(context, http.StatusOK, "group", echo.Map{
		"title": group.Name,
		"group": group,
	})
}
