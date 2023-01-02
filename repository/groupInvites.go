package repository

import (
	"photo-sharing/db"
	"photo-sharing/model"
)

func GetInvites(userId uint, dest interface{}, preloads ...string) error {
	query := db.DB.
		Model(&model.GroupInvite{}).
		Where("user_id = ?", userId)

	for _, p := range preloads {
		query = query.Preload(p)
	}

	err := query.Find(dest).Error
	return err
}
