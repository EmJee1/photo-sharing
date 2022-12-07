package model

import "time"

type User struct {
	ID           uint
	Email        string
	Password     string
	CreatedAt    time.Time
	Posts        []Post
	Groups       []*Group      `gorm:"many2many:group_users;"`
	GroupInvites []GroupInvite `gorm:"foreignKey:InvitedBy"`
}
