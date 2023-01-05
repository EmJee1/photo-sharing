package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/dto"
	"photo-sharing/model"
	"photo-sharing/repository"
	"photo-sharing/util"
	"strconv"
)

func PostPost(context echo.Context) error {
	userId := context.Get("userId").(uint)
	caption := context.FormValue("caption")
	groupId, _ := strconv.ParseUint(context.FormValue("group"), 10, 64)

	var isMemberOfGroup bool
	repository.UserIsGroupMember(userId, uint(groupId), &isMemberOfGroup)

	if !isMemberOfGroup {
		return errors.New("user is not a member of that group")
	}

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

	return context.JSON(http.StatusCreated, dto.SuccessResponse{
		Ok: true,
	})
}
