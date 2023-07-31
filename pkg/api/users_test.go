package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	data "github.com/luccasbarros/the-service/internal/postgres"
)

func TestGetUsersHandler(t *testing.T) {
	db, err := data.InitPool()
	if err != nil {
		t.Errorf("init pool error: %v", err)
	}

	repository := data.New(db)
	testedFunc := &UsersHandler{
		repository: repository,
	}

	t.Run("should return 200 and all users data", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/users?limit=10&page=1", nil)
		if err != nil {
			t.Errorf("[error creating request]: %v", err)
		}

		rr := httptest.NewRecorder()

		http.HandlerFunc(testedFunc.GetAllUsers).ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

}
