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
	err := repository.GetGroup(uint(groupId), &group, "Users", "Invites", "Posts.User")

	if err != nil {
		return context.JSON(http.StatusNotFound, dto.ErrorResponse{
			Ok:    false,
			Error: "De groep bestaat niet",
		})
	}

	var alreadyInvited, alreadyInGroup bool
	for _, inv := range group.Invites {
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

	invite := model.Invite{GroupID: uint(groupId), UserID: uint(user.ID), InvitedByID: invitedById}
	db.DB.Create(&invite)

	return context.JSON(http.StatusOK, dto.SuccessResponse{
		Ok: true,
	})
}

func GetInvites(context echo.Context) error {
	userId := context.Get("userId").(uint)
	var invites []model.Invite
	repository.GetInvites(userId, &invites, "Group")

	return context.JSON(http.StatusOK, dto.GetInvitesSuccessResponse{
		SuccessResponse: dto.SuccessResponse{Ok: true},
		Invites:         invites,
	})
}

func PostInviteRespond(context echo.Context) error {
	inviteId, _ := strconv.ParseUint(context.FormValue("inviteId"), 10, 64)
	accepted := context.FormValue("action") == "accept"
	userId := context.Get("userId").(uint)

	invite := &model.Invite{}
	repository.GetInvite(uint(inviteId), &invite)

	if userId != invite.UserID {
		return context.JSON(http.StatusForbidden, dto.ErrorResponse{
			Ok:    false,
			Error: "Die uitnodiging is niet aan jou verstuurd",
		})
	}

	if accepted {
		repository.AddUserToGroup(invite.GroupID, userId)
	}

	repository.DeleteInvite(uint(inviteId))

	return context.JSON(http.StatusOK, dto.SuccessResponse{
		Ok: true,
	})
}
