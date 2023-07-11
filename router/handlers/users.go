package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/luccasbarros/the-service/internal/data"
	"github.com/luccasbarros/the-service/internal/dto"
	"github.com/luccasbarros/the-service/pkg/errors"
	req "github.com/luccasbarros/the-service/pkg/http"
)

type UsersHandler struct {
	repository UserRepository
}

type UserRepository interface {
	GetAllUsers(ctx context.Context, limit, page uint64) ([]dto.User, error)
}

func NewUsersHandler(dal *data.Data) *UsersHandler {
	return &UsersHandler{
		repository: dal,
	}
}

func (u *UsersHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
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
		errors.RespondError(w, http.StatusInternalServerError, errors.InternalServerError)
	}

	req.Respond(w, users, http.StatusOK)
}
