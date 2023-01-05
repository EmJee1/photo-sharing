package model

import "time"

type Post struct {
	ID        uint
	Filepath  string
	Caption   string
	UserID    uint
	User      User
	GroupID   uint
	Likes     []User `gorm:"many2many:likes;"`
	Comments  []Comment
	CreatedAt time.Time
}
