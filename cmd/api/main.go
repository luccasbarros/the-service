package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	data "github.com/luccasbarros/the-service/internal/postgres"
	api "github.com/luccasbarros/the-service/pkg/api"
)

type API struct {
	server *http.Server
	logger *log.Logger
}

func NewAPI() *API {
	db, err := data.InitPool()
	if err != nil {
		log.Fatalf("init pool error %v", err)
	}

	dal := data.New(db)

	apiInstance := &API{
		logger: log.New(os.Stdout, "[SERVER] ", log.Ldate|log.Ltime),
	}

	apiInstance.server = &http.Server{
		Handler:           api.NewHandler(dal),
		Addr:              "localhost:8080",
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	return apiInstance
}

func main() {
	err := godotenv.Load("../../.envx")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	api := NewAPI()

	api.logger.Println("Server is running on http://localhost:8080")

	err = api.server.ListenAndServe()
	if err != nil {
		api.logger.Fatalf("Error starting the server %v ", err)
	}
}
