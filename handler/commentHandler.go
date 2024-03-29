package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/dto"
	"photo-sharing/model"
	"photo-sharing/repository"
	"photo-sharing/util"
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

	repository.CreateComment(&model.Comment{
		UserID: userId,
		Body:   body,
		PostID: uint(postId),
	})

	return context.JSON(http.StatusCreated, dto.SuccessResponse{
		Ok: true,
	})
}

func DeleteComment(context echo.Context) error {
	userId := context.Get("userId").(uint)
	commentId, _ := strconv.ParseUint(context.FormValue("commentId"), 10, 64)

	var userIsAdminIn []uint
	comment := &model.Comment{}
	repository.GetComment(uint(commentId), &comment, "Post")
	repository.GetGroupsUserIsAdminOf(userId, &userIsAdminIn)

	isAdminUser := util.Contains(userIsAdminIn, comment.Post.GroupID)

	if !(isAdminUser || comment.UserID == userId) {
		return context.JSON(http.StatusForbidden, dto.ErrorResponse{
			Ok:    false,
			Error: "Je hebt geen rechten om die comment te verwijderen",
		})
	}

	repository.DeleteComment(uint(commentId))
	return context.JSON(http.StatusOK, dto.SuccessResponse{
		Ok: true,
	})
}
