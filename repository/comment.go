package repository

import (
	"photo-sharing/db"
	"photo-sharing/model"
)

func GetComment(commentId uint, dest interface{}, preloads ...string) error {
	query := db.DB.Model(&model.Comment{})

	for _, p := range preloads {
		query = query.Preload(p)
	}

	return query.First(dest, commentId).Error
}
