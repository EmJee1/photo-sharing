package model

type GroupInvite struct {
	ID uint
	// This does not have a gorm reference because we want to be able to
	// invite people that might sign up later
	InviteeEmail string
	InvitedBy    uint
	GroupID      uint
}
