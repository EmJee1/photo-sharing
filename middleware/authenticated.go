package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/util"
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

		context.Set("userId", tokenContent)

		return next(context)
	}
}
