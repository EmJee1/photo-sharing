package util

import (
	"github.com/foolin/goview"
	"html/template"
	"photo-sharing/model"
)

var GoviewConfig = goview.Config{
	Root:         "views",
	Extension:    ".html",
	Master:       "layouts/master",
	Partials:     []string{"partials/post", "partials/userIcon", "partials/comment"},
	DisableCache: true,
	Funcs: template.FuncMap{
		"userLikedPost": func(userId uint, post model.Post) bool {
			for _, u := range post.Likes {
				if u.ID == userId {
					return true
				}
			}
			return false
		},
		"canDeleteComment": func(comment model.Comment, user model.User) bool {
			for _, groupId := range user.IsAdminIn {
				if groupId == comment.Post.GroupID {
					return true
				}
			}
			return comment.UserID == user.ID
		},
		"canDeletePost": func(post model.Post, user model.User) bool {
			for _, groupId := range user.IsAdminIn {
				if groupId == post.GroupID {
					return true
				}
			}
			return post.UserID == user.ID
		},
		"canKickUser": func(userToKick, user model.User, groupId uint) bool {
			userToKickIsAdmin := Contains(userToKick.IsAdminIn, groupId)
			isAdmin := Contains(user.IsAdminIn, groupId)
			return !userToKickIsAdmin && isAdmin
		},
		"canPromoteUser": func(userToPromote, user model.User, groupId uint) bool {
			userToPromoteIsAdmin := Contains(userToPromote.IsAdminIn, groupId)
			isAdmin := Contains(user.IsAdminIn, groupId)
			return !userToPromoteIsAdmin && isAdmin
		},
		"containsUint": func(slice []uint, el uint) bool {
			return Contains(slice, el)
		},
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key := values[i].(string)
				dict[key] = values[i+1]
			}
			return dict, nil
		},
	},
}
