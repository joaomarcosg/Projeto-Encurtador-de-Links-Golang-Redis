package main

import (
	"encurtador-links/internal/api"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {

	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		os.Exit(1)
	}

	slog.Info("all system offline")

}

func run() error {

	handler := api.NewHandler()

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil

}
