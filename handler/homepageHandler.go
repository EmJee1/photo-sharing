package handler

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
	"net/http"
	"photo-sharing/model"
	"photo-sharing/repository"
)

func GetHomepage(context echo.Context) error {
	userId := context.Get("userId").(uint)

	user := &model.User{}
	repository.GetUser(
		userId,
		&user,
		"Groups.Posts.Comments.User",
		"Groups.Posts."+clause.Associations,
	)

	var feed []model.Post
	for _, group := range user.Groups {
		for _, post := range group.Posts {
			feed = append(feed, post)
		}
	}

	return echoview.Render(context, http.StatusOK, "homepage", echo.Map{
		"title": "Overview",
		"user":  user,
		"feed":  feed,
	})
}
