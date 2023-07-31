package api

import data "github.com/luccasbarros/the-service/internal/postgres"

type AppHandler struct {
	UsersHandler *UsersHandler
	// others handlers
}

func NewAppHandler(dal *data.Data) *AppHandler {
	return &AppHandler{
		UsersHandler: NewUsersHandler(dal),
	}
}
