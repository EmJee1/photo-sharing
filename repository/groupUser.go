package repository

import (
	"photo-sharing/db"
	"photo-sharing/model"
)

func GetGroupsUserIsAdminOf(userId uint, dest *[]uint) error {
	var groupUsers []model.GroupUser
	err := db.DB.Where("user_id = ? AND is_admin = ?", userId, true).Find(&groupUsers).Error

	for _, groupUser := range groupUsers {
		*dest = append(*dest, groupUser.GroupID)
	}

	return err
}
