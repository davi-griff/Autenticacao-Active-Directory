package interfaces

import "auth-ad/src/internal/models"

type IApiRepository interface {
	GetRequest() ([]models.AuthRequest, error)
	SendResponse(requestId string, response models.AuthResponse) error
}

type IApiService interface {
	GetRequest() ([]models.AuthRequest, error)
	SendResponse(requestID string, response models.AuthResponse) error
}
