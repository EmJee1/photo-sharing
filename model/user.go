package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Username  string `gorm:"uniqueIndex"`
	CreatedAt time.Time
	Posts     []Post
	Likes     []Post   `gorm:"many2many:likes;"`
	Groups    []Group  `gorm:"many2many:group_users;"`
	Invites   []Invite `gorm:"foreignKey:InvitedByID"`
	Comments  []Comment
	IsAdminIn []uint `gorm:"-"`
}

func (u *User) AfterFind(tx *gorm.DB) error {
	var groupUsers []GroupUser
	tx.Where("user_id = ? AND is_admin = ?", u.ID, true).Find(&groupUsers)
	var isAdminIn []uint
	for _, user := range groupUsers {
		isAdminIn = append(isAdminIn, user.GroupID)
	}
	u.IsAdminIn = isAdminIn
	return nil
}
