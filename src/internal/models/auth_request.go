package models

type AuthRequest struct {
	RequestID string `json:"request_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
