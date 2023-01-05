package repository

import (
	"photo-sharing/db"
	"photo-sharing/model"
)

func DeleteLike(postId uint, userId uint) error {
	return db.DB.
		Model(&model.Post{ID: postId}).
		Association("Likes").
		Delete(&model.User{ID: userId})
}

func CreateLike(postId uint, userId uint) error {
	return db.DB.
		Model(&model.User{ID: userId}).
		Association("Likes").
		Append(&model.Post{ID: postId})
}
