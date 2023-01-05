package model

import "time"

type Comment struct {
	ID        uint
	Body      string
	UserID    uint
	User      User
	PostID    uint
	Post      Post
	CreatedAt time.Time
}
