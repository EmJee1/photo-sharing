package repository

import (
	"photo-sharing/db"
	"photo-sharing/model"
)

func GetUser(userId string, dest interface{}, preloads ...string) error {
	return db.DB.
		Model(&model.User{}).
		Where("id = ?", userId).
		Find(dest).
		Error
}

func GetUserByEmail(email string, dest interface{}) error {
	return db.DB.
		Model(&model.User{}).
		Where("email = ?", email).
		First(dest).
		Error
}
