package handler

import (
	"fmt"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/model"
	"strconv"
)

func PostGroupSendInvite(context echo.Context) error {
	// TODO: can only update if logged-in-user is member of group

	groupId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	inviteeEmail := context.FormValue("email")
	inviterId := context.Get("userId").(uint)

	group := &model.Group{}
	err = db.DB.
		Model(&model.Group{}).
		Where("id = ?", groupId).
		Preload("Users").
		First(&group).
		Error

	if err != nil {
		// TODO: 404 page: group not found
	}

	var exists bool
	db.DB.
		Model(&model.GroupInvite{}).
		Select("count(*) > 0").
		Where("group_id = ? AND invitee_email = ?", groupId, inviteeEmail).
		Find(&exists)

	if exists {
		// TODO: error message: invite already sent by ...
		return echoview.Render(context, http.StatusConflict, "group", echo.Map{
			"title": group.Name,
		})
	}

	groupInvite := model.GroupInvite{InviteeEmail: inviteeEmail, InvitedBy: inviterId, GroupID: uint(groupId)}
	db.DB.Create(&groupInvite)

	return context.Redirect(http.StatusSeeOther, fmt.Sprintf("/group/%d", groupId))
}
