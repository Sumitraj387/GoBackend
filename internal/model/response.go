package model

type UserResponse struct {
	Id       int64  `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone_number"`
	Email    string `json:"email_id"`
	IsActive bool   `json:"is_active"`
}
