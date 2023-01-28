package handler

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
	"net/http"
	"photo-sharing/model"
	"photo-sharing/repository"
	"sort"
)

func GetHomepage(context echo.Context) error {
	userId := context.Get("userId").(uint)

	user := &model.User{}
	repository.GetUser(
		userId,
		&user,
		"Groups.Posts.Comments."+clause.Associations,
		"Groups.Posts."+clause.Associations,
	)

	var feed []model.Post
	for _, group := range user.Groups {
		for _, post := range group.Posts {
			feed = append(feed, post)
		}
	}

	sort.Slice(feed, func(i, j int) bool {
		return feed[i].CreatedAt.After(feed[j].CreatedAt)
	})

	return echoview.Render(context, http.StatusOK, "homepage", echo.Map{
		"title": "Overview",
		"user":  user,
		"feed":  feed,
	})
}
