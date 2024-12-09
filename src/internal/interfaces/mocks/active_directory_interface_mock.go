package mocks

import (
	"auth-ad/src/internal/models"

	"github.com/stretchr/testify/mock"
)

// IActiveDirectoryInterface é um mock para a interface IActiveDirectoryInterface
type IActiveDirectoryInterface struct {
	mock.Mock
}

// Authenticate é um mock para o método Authenticate
func (m *IActiveDirectoryInterface) Authenticate(username, password string) (bool, error) {
	args := m.Called(username, password)
	return args.Bool(0), args.Error(1)
}

// GetUser é um mock para o método GetUser
func (m *IActiveDirectoryInterface) GetUser(username string) (*models.ADUser, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.ADUser), args.Error(1)
}

// GetUsers é um mock para o método GetUsers
func (m *IActiveDirectoryInterface) GetUsers(group string) ([]*models.ADUser, error) {
	args := m.Called(group)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.ADUser), args.Error(1)
}

// Bind é um mock para o método Bind
func (m *IActiveDirectoryInterface) Bind(username, password string) error {
	args := m.Called(username, password)
	return args.Error(0)
}

// Unbind é um mock para o método Unbind
func (m *IActiveDirectoryInterface) Unbind() error {
	args := m.Called()
	return args.Error(0)
}

// Close é um mock para o método Close
func (m *IActiveDirectoryInterface) Close() error {
	args := m.Called()
	return args.Error(0)
}
