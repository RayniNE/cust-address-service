package repos

import (
	"database/sql"
	"fmt"

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

func (u *UserRepo) GetAllUsers() ([]*models.User, error) {

	users := make([]*models.User, 0)

	rows, err := u.DB.Query("SELECT id, name, last_name FROM user")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user := &models.User{}
		err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Lastname,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	for _, user := range users {

		addresses, err := u.GetUserAddresses(user.Id)
		if err != nil {
			return nil, err
		}

		user.Addresses = addresses
	}

	return users, nil
}
func (u *UserRepo) GetUserById(userId int64) (*models.User, error) {

	user := &models.User{}

	err := u.DB.QueryRow("SELECT id, name, last_name FROM user WHERE id = ?", userId).Scan(&user.Id, &user.Name, &user.Lastname)
	if err != nil {
		return nil, err
	}

	addresses, err := u.GetUserAddresses(userId)
	if err != nil {
		return nil, err
	}

	user.Addresses = addresses

	return user, nil
}

func (u *UserRepo) CreateUser(user *models.User) (int64, error) {
	appContext := "UserRepo.CreateUser"

	var userId *int64

	tx, err := u.DB.Begin()
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	query := "INSERT INTO user(name, last_name) VALUES (?, ?)"
	getIdQuery := "SELECT LAST_INSERT_ID()"

	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return -1, fmt.Errorf("%v. %v", appContext, err)
	}

	_, err = stmt.Exec(user.Name, user.Lastname)
	if err != nil {
		tx.Rollback()
		return -1, fmt.Errorf("%v. %v", appContext, err)
	}

	err = tx.QueryRow(getIdQuery).Scan(&userId)
	if err != nil {
		tx.Rollback()
		return -1, fmt.Errorf("%v. %v", appContext, err)
	}

	if len(user.Addresses) > 0 {
		err = u.addPatientAddresses(tx, *userId, user.Addresses)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return -1, fmt.Errorf("%v. %v", appContext, err)
	}

	return *userId, nil
}

func (u *UserRepo) GetUserAddresses(userId int64) ([]*models.Address, error) {
	appContext := "UserRepo.GetUserAddresses"
	query := "SELECT id, user_id, address FROM address WHERE user_id = ?"

	addresses := make([]*models.Address, 0)

	rows, err := u.DB.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("%v. %v", appContext, err)
	}

	defer rows.Close()

	for rows.Next() {
		address := &models.Address{}
		err = rows.Scan(
			&address.Id,
			&address.UserId,
			&address.Address,
		)
		if err != nil {
			return nil, fmt.Errorf("%v. %v", appContext, err)
		}

		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (u *UserRepo) UpdateUser(user *models.User) error {
	appContext := "UserRepo.UpdateUser"

	tx, err := u.DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	query := "UPDATE user SET name = ?, last_name = ? WHERE id = ?"

	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("%v. %v", appContext, err)
	}

	_, err = stmt.Exec(user.Name, user.Lastname, user.Id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("%v. %v", appContext, err)
	}

	if len(user.Addresses) > 0 {
		err = u.updatePatientAddress(tx, user.Addresses)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("%v. %v", appContext, err)
	}
	return nil
}

func (u *UserRepo) addPatientAddresses(tx *sql.Tx, userId int64, addresses []*models.Address) error {
	appContext := "UserRepo.addPatientAddresses"
	query := "INSERT INTO address(user_id, address) VALUES (?, ?)"

	for _, address := range addresses {
		_, err := tx.Exec(query, userId, address.Address)
		if err != nil {
			return fmt.Errorf("%v. %v", appContext, err)
		}
	}

	return nil
}

func (u *UserRepo) updatePatientAddress(tx *sql.Tx, addresses []*models.Address) error {
	appContext := "UserRepo.updatePatientAddress"
	query := "UPDATE address SET address = ? WHERE id = ?"

	for _, address := range addresses {
		_, err := tx.Exec(query, address.Address, address.Id)
		if err != nil {
			return fmt.Errorf("%v. %v", appContext, err)
		}
	}

	return nil
}
