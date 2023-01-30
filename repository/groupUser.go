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

func DeleteGroupUser(userId, groupId uint) error {
	return connection().
		Where("user_id = ? AND group_id = ?", userId, groupId).
		Delete(&model.GroupUser{}).
		Error
}

func UpdateGroupUserAdminStatus(userId, groupId uint, isAdmin bool) error {
	return connection().
		Model(&model.GroupUser{}).
		Where("user_id = ? AND group_id = ?", userId, groupId).
		Update("is_admin", isAdmin).
		Error
}
