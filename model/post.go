package model

import "time"

type Post struct {
	ID        uint
	Filepath  string
	Caption   string
	UserID    uint
	User      User
	GroupID   uint
	CreatedAt time.Time
}
