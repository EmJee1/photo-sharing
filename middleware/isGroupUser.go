package middleware

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/repository"
	"strconv"
)

func IsGroupUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		userId := context.Get("userId").(uint)
		groupId, parseErr := strconv.ParseUint(context.Param("id"), 10, 64)

		var isGroupUser bool
		repository.UserIsGroupMember(userId, uint(groupId), &isGroupUser)

		if !isGroupUser || parseErr != nil {
			return echoview.Render(context, http.StatusNotFound, "404", echo.Map{
				"title": "Not Found",
			})
		}

		return next(context)
	}
}
