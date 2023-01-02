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

	return query.Find(dest).Error
}

func GetUserByEmail(email string, dest interface{}) error {
	return db.DB.
		Model(&model.User{}).
		Where("email = ?", email).
		First(dest).
		Error
}
