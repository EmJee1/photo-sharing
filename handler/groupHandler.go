package handler

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/model"
	"photo-sharing/repository"
	"strconv"
)

func GetGroup(context echo.Context) error {
	groupId, _ := strconv.ParseUint(context.Param("id"), 10, 64)

	group := &model.Group{}
	err := repository.GetGroup(uint(groupId), &group, "Users", "Invites.User", "Posts.User")

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

func PostGroup(context echo.Context) error {
	name := context.FormValue("name")
	description := context.FormValue("description")

	userId := context.Get("userId").(uint)

	group := model.Group{Name: name, Description: description}
	db.DB.Create(&group).Association("Users").Append([]*model.User{{ID: userId}})

	context.Redirect(http.StatusSeeOther, "/")
	return nil
}
