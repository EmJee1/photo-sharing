package repository

import (
	"photo-sharing/db"
	"photo-sharing/model"
)

func GetGroup(groupId uint, dest interface{}, preloads ...string) error {
	query := db.DB.
		Model(&model.Group{}).
		Where("id = ?", groupId)

	for _, p := range preloads {
		query = query.Preload(p)
	}

	return query.First(dest).Error
}
