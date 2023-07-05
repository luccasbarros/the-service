package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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
	conf   *Config
	server *http.Server
	logger *log.Logger
}

func NewApi(conf *Config) *API {
	// db := postgres.NewPool(conf.Database.URI)

	api := &API{
		conf:   conf,
		logger: log.New(os.Stdout, "[SERVER] ", log.Ldate|log.Ltime),
	}

	api.server = &http.Server{
		Handler:           api.Routes(),
		Addr:              "localhost:8080",
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return nil
}

func main() {
	conf := NewConfig()
	api := NewApi(conf)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(api.server.Addr, api.server.Handler)

	// Define a request handler function

	// srv.logger.Println("Server is running")

	// log.Fatal(srv.httpServer.ListenAndServe())
}
