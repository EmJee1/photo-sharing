package handler

import (
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
	groupId, _ := strconv.ParseUint(context.FormValue("groupId"), 10, 64)

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

func DeletePost(context echo.Context) error {
	userId := context.Get("userId").(uint)
	postId, _ := strconv.ParseUint(context.FormValue("postId"), 10, 64)

	post := &model.Post{}
	repository.GetPost(uint(postId), &post)

	if post.UserID != userId {
		return context.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Ok:    false,
			Error: "Je hebt geen rechten om die post te verwijderen",
		})
	}

	db.DB.Delete(model.Post{ID: uint(postId)})

	return context.JSON(http.StatusOK, dto.SuccessResponse{
		Ok: true,
	})
}
