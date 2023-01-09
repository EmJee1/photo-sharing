package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/dto"
	"photo-sharing/model"
	"photo-sharing/repository"
	"strconv"
)

// PostLike This handler toggles the like status
func PostLike(context echo.Context) error {
	userId := context.Get("userId").(uint)
	postId, _ := strconv.ParseUint(context.FormValue("postId"), 10, 64)

	post := &model.Post{}
	repository.GetPost(uint(postId), &post, "Likes")

	var isGroupUser bool
	repository.UserIsGroupMember(userId, post.GroupID, &isGroupUser)

	if !isGroupUser {
		return context.JSON(http.StatusForbidden, dto.ErrorResponse{
			Ok:    false,
			Error: "Je hebt geen rechten om die post te liken",
		})
	}

	var userLikedPost bool
	for _, l := range post.Likes {
		if l.ID == userId {
			userLikedPost = true
		}
	}

	if userLikedPost {
		repository.DeleteLike(uint(postId), userId)
	} else {
		repository.CreateLike(uint(postId), userId)
	}

	return context.JSON(http.StatusCreated, dto.SuccessResponse{
		Ok: true,
	})
}
