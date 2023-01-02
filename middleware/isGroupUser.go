package middleware

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
)

func IsGroupUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		userId := context.Get("userId").(uint)
		groupId := context.Param("id")

		var isGroupUser bool
		db.DB.
			Select("count(*) > 0").
			Table("group_users").
			Where("user_id = ? AND group_id = ?", userId, groupId).
			Find(&isGroupUser)

		if !isGroupUser {
			return echoview.Render(context, http.StatusNotFound, "404", echo.Map{
				"title": "Not Found",
			})
		}

		return next(context)
	}
}
