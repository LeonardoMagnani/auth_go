package models

type VerifyUserRequest struct {
	Code int `json:"code" binding:"required"`
}
