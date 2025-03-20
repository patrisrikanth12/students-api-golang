package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/patrisrikanth12/students-api-golang/internal/config"
	"github.com/patrisrikanth12/students-api-golang/internal/http/handlers/student"
	"github.com/patrisrikanth12/students-api-golang/internal/storage/sqlite"
)

func main() {
	// 1. Load config
	cfg := config.MustLoad()

	// 2. Load database

	storage, err := sqlite.New(cfg); 
	if err != nil {
		log.Fatal("Unable to setup the DB ", err.Error())
	}

	slog.Info("Successfully setup DB")

	// 3. Setup Router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.Create(storage))

	// 4. Setup Server
	server := http.Server{
		Addr: cfg.HttpServer.Address,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.HttpServer.Address))

	done := make(chan  os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Unable to start server")
		}
	} ()


	<- done	

	slog.Info("Shutting down the server");

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")
}