package repos

import (
	"database/sql"

	"github.com/raynine/customeraddresses/models"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (u *UserRepo) GetUserById(userId int64) (*models.User, error) {

	user := &models.User{}

	err := u.DB.QueryRow("SELECT id, name, last_name FROM user WHERE id = ?", userId).Scan(&user.Id, &user.Name, &user.Lastname)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) CreateUser(user *models.User) error {

	return nil
}

func (u *UserRepo) GetUserAddresses(userId int64) ([]*models.Address, error) {

	return []*models.Address{}, nil
}

func (u *UserRepo) UpdateUser(user *models.User) error {

	return nil
}
