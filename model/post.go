package model

import "time"

type Post struct {
	ID        uint
	Filepath  string
	Caption   string
	UserID    uint
	User      User
	GroupID   uint
	Likes     []Post `gorm:"many2many:likes;"`
	CreatedAt time.Time
}
