package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/raynine/customeraddresses/helpers"
	"github.com/raynine/customeraddresses/interfaces"
	"github.com/raynine/customeraddresses/repos"
)

type Handler struct {
	DB       *sql.DB
	UserRepo interfaces.UsersRepo
}

func CreateNewHandler(db *sql.DB) *Handler {
	return &Handler{
		DB:       db,
		UserRepo: repos.NewUserRepo(db),
	}
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		helpers.EncodeResponse(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.UserRepo.GetUserById(userId)
	if err != nil {
		helpers.EncodeResponse(w, http.StatusInternalServerError, err)
		return
	}

	helpers.EncodeResponse(w, http.StatusOK, response)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetUserAddresses(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}
