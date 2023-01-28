package repository

import (
	"photo-sharing/model"
)

func GetComment(commentId uint, dest interface{}, preloads ...string) error {
	query := connection().Model(&model.Comment{})

	for _, p := range preloads {
		query = query.Preload(p)
	}

	return query.First(dest, commentId).Error
}

func CreateComment(comment *model.Comment) error {
	return connection().Create(comment).Error
}

func DeleteComment(commentId uint) error {
	return connection().Delete(&model.Comment{}, commentId).Error
}
