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

func PostKick(context echo.Context) error {
	groupId, _ := strconv.ParseUint(context.FormValue("groupId"), 10, 64)
	userIdToKick, _ := strconv.ParseUint(context.FormValue("userId"), 10, 64)

	userToDelete := &model.User{}
	repository.GetUser(uint(userIdToKick), &userToDelete)

	if util.Contains(userToDelete.IsAdminIn, uint(groupId)) {
		return context.JSON(http.StatusForbidden, dto.ErrorResponse{
			Ok:    false,
			Error: "Je kunt geen beheerders verwijderen uit de groep",
		})
	}

	repository.DeleteGroupUser(uint(userIdToKick), uint(groupId))
	return context.JSON(http.StatusOK, dto.SuccessResponse{
		Ok: true,
	})
}
