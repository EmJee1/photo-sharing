package model

type GroupInvite struct {
	ID          uint
	UserID      uint
	User        User
	InvitedByID uint
	InvitedBy   User
	GroupID     uint
	Group       Group
}
