package repository

import (
	"photo-sharing/model"
)

func GetUser(userId uint, dest interface{}, preloads ...string) error {
	query := connection().Model(&model.User{})

	for _, p := range preloads {
		query = query.Preload(p)
	}

	return query.First(dest, userId).Error
}

func GetUserByEmail(email string, dest interface{}) error {
	err := connection().
		Model(&model.User{}).
		Where("email = ?", email).
		First(dest).
		Error
	return err
}

func UserIsGroupMember(userId uint, groupId uint, dest interface{}) error {
	return connection().
		Select("count(*) > 0").
		Table("group_users").
		Where("user_id = ? AND group_id = ?", userId, groupId).
		Find(dest).
		Error
}

func CreateUser(user *model.User) error {
	return connection().Create(user).Error
}
