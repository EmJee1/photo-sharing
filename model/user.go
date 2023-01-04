package model

import "time"

type User struct {
	ID        uint
	Email     string
	Password  string
	Username  string
	CreatedAt time.Time
	Posts     []Post
	Likes     []Post   `gorm:"many2many:likes;"`
	Groups    []Group  `gorm:"many2many:group_users;"`
	Invites   []Invite `gorm:"foreignKey:InvitedByID"`
}
