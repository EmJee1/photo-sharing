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

func AddUserToGroup(groupId uint, userId uint) error {
	err := db.DB.
		Model(&model.Group{ID: groupId}).
		Association("Users").
		Append(&model.User{ID: userId})
	return err
}
