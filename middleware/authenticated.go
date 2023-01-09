package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/dto"
	"photo-sharing/util"
	"strconv"
)

func unauthenticatedResponse(context echo.Context) {
	if util.IsApiRoute(context) {
		context.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Ok:    false,
			Error: "Je moet zijn ingelogd om dit te kunnen doen",
		})
		return
	}

	context.Redirect(http.StatusFound, "/login")
}

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		cookie, err := context.Cookie("token")
		if err != nil {
			unauthenticatedResponse(context)
			return nil
		}

		tokenContent, err := util.ParseJWT(cookie.Value)
		if err != nil {
			unauthenticatedResponse(context)
			return nil
		}

		userId, err := strconv.ParseUint(tokenContent, 10, 64)
		if err != nil {
			unauthenticatedResponse(context)
			return nil
		}

		context.Set("userId", uint(userId))

		return next(context)
	}
}
