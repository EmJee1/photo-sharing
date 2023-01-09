package util

import (
	"github.com/labstack/echo/v4"
	"strings"
)

func IsApiRoute(context echo.Context) bool {
	return strings.HasPrefix(context.Request().URL.Path, "/api")
}
