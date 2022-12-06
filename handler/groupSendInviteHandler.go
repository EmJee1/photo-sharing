package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/dto"
	"photo-sharing/model"
	"strconv"
)

func PostGroupSendInvite(context echo.Context) error {
	// TODO: can only invite if logged-in-user is member of group

	groupId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	inviteeEmail := context.FormValue("inviteeEmail")
	inviterId := context.Get("userId").(uint)

	group := &model.Group{}
	err = db.DB.
		Model(&model.Group{}).
		Where("id = ?", groupId).
		Preload("Users").
		First(&group).
		Error

	if err != nil {
		return context.JSON(http.StatusConflict, dto.ErrorResponse{
			Error: "That group does not exist",
			Ok:    false,
		})
	}

	var exists bool
	db.DB.
		Model(&model.GroupInvite{}).
		Select("count(*) > 0").
		Where("group_id = ? AND invitee_email = ?", groupId, inviteeEmail).
		Find(&exists)

	if exists {
		return context.JSON(http.StatusConflict, dto.ErrorResponse{
			Error: "That user is already invited for this group",
			Ok:    false,
		})
	}

	groupInvite := model.GroupInvite{InviteeEmail: inviteeEmail, InvitedBy: inviterId, GroupID: uint(groupId)}
	db.DB.Create(&groupInvite)

	return context.JSON(http.StatusCreated, dto.SuccessResponse{
		Ok: true,
	})
}
