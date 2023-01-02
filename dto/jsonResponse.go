package dto

import "photo-sharing/model"

type ErrorResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

type SuccessResponse struct {
	Ok bool `json:"ok"`
}

type GetInvitesSuccessResponse struct {
	Invites []model.GroupInvite `json:"invites"`
	SuccessResponse
}
