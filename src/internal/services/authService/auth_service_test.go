package authService

import (
	"testing"

	"auth-ad/src/internal/interfaces/mocks"
	"auth-ad/src/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	mockRepo := new(mocks.IActiveDirectoryInterface)
	service := NewAuthService(mockRepo)

	mockRepo.On("Authenticate", "user", "pass").Return(true, nil)

	authenticated, err := service.Authenticate("user", "pass")
	assert.NoError(t, err)
	assert.True(t, authenticated)
}

func TestGetUser(t *testing.T) {
	mockRepo := new(mocks.IActiveDirectoryInterface)
	service := NewAuthService(mockRepo)

	mockADUser := &models.ADUser{
		SAMAccountName: "user",
		Email:          "user@example.com",
		Groups:         []string{"group1"},
	}

	mockRepo.On("GetUser", "user").Return(mockADUser, nil)

	user, err := service.GetUser("user")
	assert.NoError(t, err)
	assert.Equal(t, "user", user.Username)
	assert.Equal(t, "user@example.com", user.Email)
}

func TestGetUsers(t *testing.T) {
	mockRepo := new(mocks.IActiveDirectoryInterface)
	service := NewAuthService(mockRepo)

	mockADUsers := []*models.ADUser{
		{
			SAMAccountName: "user1",
			Email:          "user1@example.com",
			Groups:         []string{"group1"},
		},
		{
			SAMAccountName: "user2",
			Email:          "user2@example.com",
			Groups:         []string{"group1"},
		},
	}

	mockRepo.On("GetUsers", "group1").Return(mockADUsers, nil)

	users, err := service.GetUsers("group1")
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}
