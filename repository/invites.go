package repository

import (
	"photo-sharing/model"
)

func GetInvite(inviteId uint, dest interface{}) error {
	return connection().
		Model(&model.Invite{}).
		First(dest, inviteId).
		Error
}

func GetInvites(userId uint, dest interface{}, preloads ...string) error {
	query := connection().
		Model(&model.Invite{}).
		Where("user_id = ?", userId)

	for _, p := range preloads {
		query = query.Preload(p)
	}

	return query.Find(dest).Error
}

func CreateInvite(invite *model.Invite) error {
	return connection().Create(invite).Error
}

func DeleteInvite(inviteId uint) error {
	return connection().Delete(&model.Invite{}, inviteId).Error
}
