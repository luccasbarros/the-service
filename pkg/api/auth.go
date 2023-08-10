package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/luccasbarros/the-service/pkg/auth"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "error decoding")
		return
	}

	token, err := auth.GenerateToken(credentials.Username)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(5 * time.Minute),
	})
}
