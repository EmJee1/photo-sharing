package repository

import (
	"photo-sharing/db"
	"photo-sharing/model"
)

func GetPost(postId uint, dest interface{}, preloads ...string) error {
	query := db.DB.Model(model.Post{})

	for _, p := range preloads {
		query = query.Preload(p)
	}

	return query.First(dest, postId).Error
}
