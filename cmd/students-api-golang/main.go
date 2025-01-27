package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/patrisrikanth12/students-api-golang/internal/config"
)

func main() {
	// 1. Load config
	cfg := config.MustLoad()

	// 2. Load database

	// 3. Setup Router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to student API"))
	})

	// 4. Setup Server
	server := http.Server{
		Addr: cfg.HttpServer.Address,
		Handler: router,
	}

	fmt.Printf("Server started: %s", cfg.HttpServer.Address)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Unable to start server")
	}
}