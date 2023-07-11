package http_handle

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}