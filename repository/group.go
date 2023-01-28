package repository

import (
	"photo-sharing/model"
)

func GetGroup(groupId uint, dest interface{}, preloads ...string) error {
	query := connection().Model(&model.Group{})

	for _, p := range preloads {
		query = query.Preload(p)
	}

	return query.First(dest, groupId).Error
}

func AddUserToGroup(groupId uint, userId uint) error {
	return connection().
		Model(&model.Group{ID: groupId}).
		Where("id = ?", groupId).
		Association("Users").
		Append(&model.User{ID: userId})
}

func CreateGroup(group *model.Group) error {
	return connection().Create(group).Error
}
