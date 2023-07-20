package handlers

import (
	"github.com/luccasbarros/the-service/internal/data"
	"github.com/luccasbarros/the-service/router/handlers/users"
)

type AppHandler struct {
	UsersHandler *users.UsersHandler
	// others handlers
}

func NewAppHandler(dal *data.Data) *AppHandler {
	return &AppHandler{
		UsersHandler: users.NewUsersHandler(dal),
	}
}