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

func DeleteComment(context echo.Context) error {
	userId := context.Get("userId").(uint)
	commentId, _ := strconv.ParseUint(context.FormValue("commentId"), 10, 64)

	comment := &model.Comment{}
	repository.GetComment(uint(commentId), &comment, "Post")

	var isGroupUser bool
	repository.UserIsGroupMember(userId, comment.Post.GroupID, &isGroupUser)

	if !isGroupUser || comment.UserID != userId {
		return context.JSON(http.StatusForbidden, dto.ErrorResponse{
			Ok:    false,
			Error: "Je hebt geen rechten om die comment te verwijderen",
		})
	}

	db.DB.Delete(&model.Comment{ID: uint(commentId)})
	return context.JSON(http.StatusOK, dto.SuccessResponse{
		Ok: true,
	})
}
