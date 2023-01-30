package middleware

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/dto"
	"photo-sharing/model"
	"photo-sharing/repository"
	"photo-sharing/util"
)

func IsGroupAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		userId := context.Get("userId").(uint)
		groupId, err := getGroupId(context)

		user := &model.User{}
		repository.GetUser(userId, &user)

		if !util.Contains(user.IsAdminIn, uint(groupId)) || err != nil {
			// Send response in JSON if the request is an API request
			if util.IsApiRoute(context) {
				return context.JSON(http.StatusNotFound, dto.ErrorResponse{
					Ok:    false,
					Error: "Je bent geen beheerder in die groep",
				})
			}

			return echoview.Render(context, http.StatusNotFound, "404", echo.Map{
				"title": "Not Found",
			})
		}

		return next(context)
	}
}
