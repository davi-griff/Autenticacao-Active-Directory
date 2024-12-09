package interfaces

import "auth-ad/src/internal/models"

type IActiveDirectoryRepository interface {
	Authenticate(username, password string) (bool, error)
	GetUser(username string) (*models.ADUser, error)
	GetUsers(group string) ([]*models.ADUser, error)
	Bind(username, password string) error
	Unbind() error
	Close() error
}

type IActiveDirectoryService interface {
	Authenticate(username, password string) (bool, error)
	GetUser(username string) (models.UserData, error)
	Unbind() error
}
