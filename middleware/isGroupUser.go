package middleware

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/dto"
	"photo-sharing/repository"
	"strconv"
	"strings"
)

func getGroupId(context echo.Context) (uint64, error) {
	groupId := context.FormValue("groupId")
	if groupId == "" {
		groupId = context.Param("id")
	}

	return strconv.ParseUint(groupId, 10, 64)
}

func IsGroupUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		userId := context.Get("userId").(uint)
		groupId, err := getGroupId(context)

		var isGroupUser bool
		repository.UserIsGroupMember(userId, uint(groupId), &isGroupUser)

		if !isGroupUser || err != nil {
			// Send response in JSON if the request is an API request
			if strings.HasPrefix(context.Request().URL.Path, "/api") {
				return context.JSON(http.StatusNotFound, dto.ErrorResponse{
					Ok:    false,
					Error: "Je hebt geen toegang tot die groep",
				})
			}

			return echoview.Render(context, http.StatusNotFound, "404", echo.Map{
				"title": "Not Found",
			})
		}

		return next(context)
	}
}
