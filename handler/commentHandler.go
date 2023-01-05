package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/dto"
	"photo-sharing/model"
	"photo-sharing/repository"
	"strconv"
)

func PostComment(context echo.Context) error {
	userId := context.Get("userId").(uint)
	body := context.FormValue("body")
	postId, _ := strconv.ParseUint(context.FormValue("postId"), 10, 64)

	post := &model.Post{}
	repository.GetPost(uint(postId), &post)

	var isGroupUser bool
	repository.UserIsGroupMember(userId, post.GroupID, &isGroupUser)

	if !isGroupUser {
		return context.JSON(http.StatusForbidden, dto.ErrorResponse{
			Ok:    false,
			Error: "Je hebt geen rechten om een comment te plaatsen bij die post",
		})
	}

	db.DB.Create(&model.Comment{UserID: userId, Body: body, PostID: uint(postId)})

	return context.JSON(http.StatusCreated, dto.SuccessResponse{
		Ok: true,
	})
}
