package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/util"
	"strconv"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		cookie, err := context.Cookie("token")
		if err != nil {
			context.Redirect(http.StatusFound, "/login")
			return nil
		}

		tokenContent, err := util.ParseJWT(cookie.Value)
		if err != nil {
			context.Redirect(http.StatusFound, "/login")
			return nil
		}

		userId, err := strconv.ParseUint(tokenContent, 10, 64)
		if err != nil {
			context.Redirect(http.StatusFound, "/login")
			return nil
		}

		context.Set("userId", uint(userId))

		return next(context)
	}
}
