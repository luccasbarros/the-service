package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type server struct {
	logger     *log.Logger
	httpServer *http.Server
}

func main() {
	// Define a request handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	}

	srv := server{
		httpServer: &http.Server{
			Handler:           http.HandlerFunc(handler),
			Addr:              "localhost:8080",
			ReadTimeout:       10 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
		},
		logger: log.New(os.Stdout, "[SERVER] ", log.Ldate|log.Ltime),
	}

	srv.logger.Println("Server is running")

	log.Fatal(srv.httpServer.ListenAndServe())
}
