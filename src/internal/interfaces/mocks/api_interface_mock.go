package mocks

import (
	"auth-ad/src/internal/models"

	"github.com/stretchr/testify/mock"
)

type IApiRepository struct {
	mock.Mock
}

func (a *IApiRepository) GetRequest() ([]models.AuthRequest, error) {
	args := a.Called()
	return args.Get(0).([]models.AuthRequest), args.Error(1)
}

func (a *IApiRepository) SendResponse(requestId string, response models.AuthResponse) error {
	args := a.Called(requestId, response)
	return args.Error(0)
}
