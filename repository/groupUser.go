package repository

import (
	"photo-sharing/model"
)

func GetGroupsUserIsAdminOf(userId uint, dest *[]uint) error {
	var groupUsers []model.GroupUser
	err := connection().Where("user_id = ? AND is_admin = ?", userId, true).Find(&groupUsers).Error

	for _, groupUser := range groupUsers {
		*dest = append(*dest, groupUser.GroupID)
	}

	return err
}

func CreateGroupUser(groupUser *model.GroupUser) error {
	return connection().Create(&groupUser).Error
}
