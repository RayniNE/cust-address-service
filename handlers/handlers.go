package handlers

import (
	"database/sql"
	"net/http"
)

type Handler struct {
	DB *sql.DB
	// UserRepo interfaces.UsersRepo
}

func CreateNewHandler(db *sql.DB) *Handler {
	return &Handler{
		DB: db,
	}
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetUserAddresses(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}
