package apiService

import (
	"errors"
	"testing"

	"auth-ad/src/internal/interfaces/mocks"
	"auth-ad/src/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestGetRequest(t *testing.T) {
	mockRepo := new(mocks.IApiRepository)
	service := NewApiService(mockRepo)

	expectedRequests := []models.AuthRequest{{}, {}}
	mockRepo.On("GetRequest").Return(expectedRequests, nil)

	requests, err := service.GetRequest()

	assert.NoError(t, err)
	assert.Equal(t, expectedRequests, requests)
	mockRepo.AssertExpectations(t)
}

func TestSendResponse(t *testing.T) {
	mockRepo := new(mocks.IApiRepository)
	service := NewApiService(mockRepo)

	requestID := "123"
	response := models.AuthResponse{}
	mockRepo.On("SendResponse", requestID, response).Return(nil)

	err := service.SendResponse(requestID, response)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSendResponse_Error(t *testing.T) {
	mockRepo := new(mocks.IApiRepository)
	service := NewApiService(mockRepo)

	requestID := "123"
	response := models.AuthResponse{}
	mockRepo.On("SendResponse", requestID, response).Return(errors.New("some error"))

	err := service.SendResponse(requestID, response)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
