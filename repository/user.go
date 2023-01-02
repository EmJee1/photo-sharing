package repository

import (
	"photo-sharing/db"
	"photo-sharing/model"
)

func GetUser(userId uint, dest interface{}, preloads ...string) error {
	query := db.DB.
		Model(&model.User{}).
		Where("id = ?", userId)

	for _, p := range preloads {
		query = query.Preload(p)
	}

	return query.First(dest).Error
}

func GetUserByEmail(email string, dest interface{}) error {
	err := db.DB.
		Model(&model.User{}).
		Where("email = ?", email).
		First(dest).
		Error
	return err
}

func UserIsGroupMember(userId uint, groupId uint, dest interface{}) error {
	return db.DB.
		Select("count(*) > 0").
		Table("group_users").
		Where("user_id = ? AND group_id = ?", userId, groupId).
		Find(dest).
		Error
}
