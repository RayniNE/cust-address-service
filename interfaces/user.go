package interfaces

import "github.com/raynine/customeraddresses/models"

type UsersRepo interface {
	GetUserById(int64) (*models.User, error)
	CreateUser(*models.User) error
	GetUserAddresses(int64) ([]*models.Address, error)
	UpdateUser(*models.User) error
}
