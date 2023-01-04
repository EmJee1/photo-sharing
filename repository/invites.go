package repository

import (
	"photo-sharing/db"
	"photo-sharing/model"
)

func GetInvite(inviteId uint, dest interface{}) error {
	err := db.DB.
		Model(&model.Invite{}).
		Where("id = ?", inviteId).
		First(dest).
		Error
	return err
}

func GetInvites(userId uint, dest interface{}, preloads ...string) error {
	query := db.DB.
		Model(&model.Invite{}).
		Where("user_id = ?", userId)

	for _, p := range preloads {
		query = query.Preload(p)
	}

	err := query.Find(dest).Error
	return err
}

func DeleteInvite(inviteId uint) error {
	err := db.DB.Delete(&model.Invite{}, inviteId).Error
	return err
}
