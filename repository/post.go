package repository

import (
	"photo-sharing/model"
)

func GetPost(postId uint, dest interface{}, preloads ...string) error {
	query := connection().Model(model.Post{})

	for _, p := range preloads {
		query = query.Preload(p)
	}

	return query.First(dest, postId).Error
}

func CreatePost(post *model.Post) error {
	return connection().Create(post).Error
}

func DeletePost(postId uint) error {
	return connection().Delete(&model.Post{}, postId).Error
}
