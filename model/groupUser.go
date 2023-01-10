package model

type GroupUser struct {
	UserID  uint `gorm:"primaryKey"`
	GroupID uint `gorm:"primaryKey"`
	IsAdmin bool `gorm:"default:false"`
}
