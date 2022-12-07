package model

type Group struct {
	ID           uint
	Name         string
	Description  string
	Posts        []Post
	Users        []*User `gorm:"many2many:group_users;"`
	GroupInvites []GroupInvite
}
