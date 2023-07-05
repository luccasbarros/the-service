package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/luccasbarros/the-service/internal/data"
)

func (a *API) HandlerTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func (a *API) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.HandlerTest)

	return mux
}

type API struct {
	server *http.Server
	logger *log.Logger
	db *pgxpool.Pool
}

func NewApi() *API {
	db, err := data.InitPool()
	if err != nil {
		log.Fatalf("init pool error %v", err)
	}

	api := &API{
		logger: log.New(os.Stdout, "[SERVER] ", log.Ldate|log.Ltime),
		db: db,
	}

	api.server = &http.Server{
		Handler:           api.Routes(),
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
