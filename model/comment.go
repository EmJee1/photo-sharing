package model

type Comment struct {
	ID     uint
	Body   string
	UserID uint
	User   User
	PostID uint
	Post   Post
}
