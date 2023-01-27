package repository

import (
	"photo-sharing/db"
	"photo-sharing/model"
)

func GetGroupsUserIsAdminOf(userId uint) []uint {
	var groupUsers []model.GroupUser
	db.DB.
		Model(&model.GroupUser{}).
		Where("user_id = ? AND is_admin = ?", userId, true).
		Find(&groupUsers)

	var groups []uint
	for _, groupUser := range groupUsers {
		if groupUser.IsAdmin {
			groups = append(groups, groupUser.GroupID)
		}
	}

	return groups
}
