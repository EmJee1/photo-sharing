package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/model"
)

// PostGroupCreate TODO: move to groupHandler file & change route
func PostGroupCreate(context echo.Context) error {
	name := context.FormValue("name")
	description := context.FormValue("description")

	userId := context.Get("userId").(uint)

	group := model.Group{Name: name, Description: description}
	db.DB.Create(&group).Association("Users").Append([]*model.User{{ID: userId}})

	context.Redirect(http.StatusSeeOther, "/")
	return nil
}
