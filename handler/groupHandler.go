package handler

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/model"
	"photo-sharing/repository"
	"sort"
	"strconv"
)

func GetGroup(context echo.Context) error {
	groupId, _ := strconv.ParseUint(context.Param("id"), 10, 64)
	userId := context.Get("userId").(uint)

	user := &model.User{}
	repository.GetUser(userId, &user)

	group := &model.Group{}
	repository.GetGroup(uint(groupId), &group, "Users", "Invites.User", "Posts.Comments."+clause.Associations, "Posts."+clause.Associations)

	sort.Slice(group.Posts, func(i, j int) bool {
		return group.Posts[i].CreatedAt.After(group.Posts[j].CreatedAt)
	})

	return echoview.Render(context, http.StatusOK, "group", echo.Map{
		"title": group.Name,
		"group": group,
		"user":  user,
	})
}

func PostGroup(context echo.Context) error {
	name := context.FormValue("name")
	description := context.FormValue("description")
	userId := context.Get("userId").(uint)

	group := model.Group{Name: name, Description: description}
	db.DB.Create(&group)

	groupUser := model.GroupUser{UserID: userId, GroupID: group.ID, IsAdmin: true}
	db.DB.Create(&groupUser)

	context.Redirect(http.StatusSeeOther, "/")
	return nil
}
