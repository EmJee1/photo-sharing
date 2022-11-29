package handler

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/model"
	"strconv"
)

func GetGroupCreate(context echo.Context) error {
	return echoview.Render(context, http.StatusOK, "createGroup", echo.Map{
		"title": "Groep Aanmaken",
	})
}

func PostGroupCreate(context echo.Context) error {
	name := context.FormValue("name")

	userId := context.Get("userId").(string)
	uintUserId, _ := strconv.ParseUint(userId, 10, 64)

	group := model.Group{Name: name}
	db.DB.Create(&group)
	db.DB.Model(&group).Association("Users").Append([]*model.User{{ID: uint(uintUserId)}})

	context.Redirect(http.StatusSeeOther, "/")
	return nil
}
