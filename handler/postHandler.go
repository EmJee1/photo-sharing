package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/model"
	"photo-sharing/util"
	"strconv"
)

func PostPost(context echo.Context) error {
	// TODO: check user is member of group
	// TODO: only accept image formats

	userId := context.Get("userId").(uint)
	caption := context.FormValue("caption")
	groupId, _ := strconv.ParseUint(context.FormValue("group"), 10, 64)

	file, err := context.FormFile("image")
	if err != nil {
		return err
	}

	filepath, err := util.UploadImage(file)
	if err != nil {
		return err
	}

	db.DB.Create(&model.Post{
		Caption:  caption,
		UserID:   userId,
		Filepath: filepath,
		GroupID:  uint(groupId),
	})

	return context.Redirect(http.StatusSeeOther, fmt.Sprintf("/group/%d", groupId))
}
