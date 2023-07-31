package api

import (
	"context"
	"log"
	"net/http"
	"strconv"

	data "github.com/luccasbarros/the-service/internal/postgres"
)

type UsersHandler struct {
	repository UserRepository
}

type UserRepository interface {
	GetAllUsers(ctx context.Context, limit, page uint64) ([]data.User, error)
}

func NewUsersHandler(dal *data.Data) *UsersHandler {
	return &UsersHandler{
		repository: dal,
	}
}

func (u *UsersHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	limitStr := queryParams.Get("limit")
	pageStr := queryParams.Get("page")

	limit, limitErr := strconv.ParseUint(limitStr, 10, 64)
	if limitErr != nil || limit < 10 || limit > 100 {
		limit = 100
	}

	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		page = 1
	}

	users, err := u.repository.GetAllUsers(r.Context(), limit, page)
	if err != nil {
		log.Println("Error getting users: ", err.Error())
		RespondError(w, http.StatusInternalServerError, InternalServerError)
	}

	Respond(w, users, http.StatusOK)
}
