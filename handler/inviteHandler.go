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

func PostInvite(context echo.Context) error {
	groupId, _ := strconv.ParseUint(context.FormValue("groupId"), 10, 64)
	userEmail := context.FormValue("email")
	invitedById := context.Get("userId").(uint)

	user := &model.User{}
	if err := repository.GetUserByEmail(userEmail, &user); err != nil {
		return context.JSON(http.StatusNotFound, dto.ErrorResponse{
			Ok:    false,
			Error: "De uitgenodigde gebruiker bestaat niet",
		})
	}

	group := &model.Group{}
	err := repository.GetGroup(uint(groupId), &group, "Users", "GroupInvites", "Posts.User")

	if err != nil {
		return context.JSON(http.StatusNotFound, dto.ErrorResponse{
			Ok:    false,
			Error: "De groep bestaat niet",
		})
	}

	var alreadyInvited, alreadyInGroup bool
	for _, inv := range group.GroupInvites {
		if inv.UserID == uint(user.ID) {
			alreadyInvited = true
		}
	}
	for _, usr := range group.Users {
		if usr.ID == uint(user.ID) {
			alreadyInGroup = true
		}
	}

	if alreadyInvited || alreadyInGroup {
		return context.JSON(http.StatusConflict, dto.ErrorResponse{
			Ok:    false,
			Error: "Die gebruiker is al uitgenodigd, of al deelnemer van de groep",
		})
	}

	groupInvite := model.GroupInvite{GroupID: uint(groupId), UserID: uint(user.ID), InvitedByID: invitedById}
	db.DB.Create(&groupInvite)

	return context.JSON(http.StatusOK, dto.SuccessResponse{
		Ok: true,
	})
}
