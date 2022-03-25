package interfaces

import "github.com/raynine/customeraddresses/models"

type UsersRepo interface {
	GetUserById(int64) (*models.User, error)
	CreateUser(*models.User) (int64, error)
	GetUserAddresses(int64) ([]*models.Address, error)
	UpdateUser(*models.User) error
	GetAllUsers() ([]*models.User, error)
}
