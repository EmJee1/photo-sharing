package dto

type ErrorResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"'`
}

type SuccessResponse struct {
	Ok bool `json:"ok"`
}
