package handler

import (
	"github.com/labstack/echo/v4"
)

func PostPost(context echo.Context) error {
	// TODO: check user is member of group

	userId := context.Get("userId").(uint)
	caption := context.Param("caption")

	file, err := context.FormFile("image")
	if err != nil {
		return err
	}
}
