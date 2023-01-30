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

func PostPromote(context echo.Context) error {
	groupId, _ := strconv.ParseUint(context.FormValue("groupId"), 10, 64)
	userIdToPromote, _ := strconv.ParseUint(context.FormValue("userId"), 10, 64)

	userToPromote := &model.User{}
	repository.GetUser(uint(userIdToPromote), &userToPromote)

	if util.Contains(userToPromote.IsAdminIn, uint(groupId)) {
		return context.JSON(http.StatusConflict, dto.ErrorResponse{
			Ok:    false,
			Error: "Die gebruiker is al een beheerder",
		})
	}

	err := repository.UpdateGroupUserAdminStatus(uint(userIdToPromote), uint(groupId), true)
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, dto.SuccessResponse{
		Ok: true,
	})
}
