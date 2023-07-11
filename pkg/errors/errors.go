package errors

import (
	"encoding/json"
	"net/http"
)

const (
	InternalServerError = "Internal Server Error"
)

type ErrorResponse struct {
	Message    string `json:"error"`
	StatusCode int    `json:"status"`
}


func RespondError(w http.ResponseWriter, statusCode int, message string) {
	errorResponse := &ErrorResponse{
		Message:    message,
		StatusCode: statusCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errJSON, err := json.Marshal(errorResponse)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(errJSON)
	if err != nil {
		panic(err)
	}
}
