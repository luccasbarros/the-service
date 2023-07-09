package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/luccasbarros/the-service/internal/data"
	"github.com/luccasbarros/the-service/router"
)

type API struct {
	server *http.Server
	logger *log.Logger
}

func NewApi() *API {
	db, err := data.InitPool()
	if err != nil {
		log.Fatalf("init pool error %v", err)
	}

	dal := data.New(db)

	api := &API{
		logger: log.New(os.Stdout, "[SERVER] ", log.Ldate|log.Ltime),
	}

	api.server = &http.Server{
		Handler:           router.NewHandler(dal),
		Addr:              "localhost:8080",
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return api
}

func main() {
	api := NewApi()

	api.logger.Println("Server is running on http://localhost:8080")

	http.ListenAndServe(api.server.Addr, api.server.Handler)
}
