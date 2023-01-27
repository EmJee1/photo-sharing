package model

type GroupUser struct {
	UserID  uint
	GroupID uint
	IsAdmin bool `gorm:"default:false"`
}
