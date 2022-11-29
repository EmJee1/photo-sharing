package middleware

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/util"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		cookie, err := context.Cookie("token")
		if err != nil {
			// TODO: Redirect to login instead of render login
			return echoview.Render(context, http.StatusUnauthorized, "login", echo.Map{
				"error": "You need to be logged in to view that page",
				"title": "Login",
			})
		}

		tokenContent, err := util.ParseJWT(cookie.Value)
		if err != nil {
			// TODO: Redirect to login instead of render login
			return echoview.Render(context, http.StatusInternalServerError, "login", echo.Map{
				"error": "The login session is no longer valid",
				"title": "Login",
			})
		}

		context.Set("userId", tokenContent)

		return next(context)
	}
}
