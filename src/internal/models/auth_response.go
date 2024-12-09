package models

type AuthResponse struct {
	RequestID string   `json:"request_id"`
	Success   bool     `json:"success"`
	UserData  UserData `json:"user_data"`
}
