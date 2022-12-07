package model

import "time"

type Post struct {
	ID        uint
	Filepath  string
	Caption   string
	UserID    uint
	GroupID   uint
	CreatedAt time.Time
}
