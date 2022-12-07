package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/model"
	"photo-sharing/util"
	"strconv"
)

func PostPost(context echo.Context) error {
	// TODO: check user is member of group

	userId := context.Get("userId").(uint)
	caption := context.Param("caption")
	groupId, _ := strconv.ParseUint(context.Param("groupId"), 10, 64)

	file, err := context.FormFile("image")
	if err != nil {
		return err
	}

	if err := util.UploadImage(file); err != nil {
		return err
	}

	post := model.Post{Caption: caption, UserID: userId, Filepath: "/file", GroupID: uint(groupId)}
	db.DB.Create(&post)

	return context.Redirect(http.StatusSeeOther, "/")
}
