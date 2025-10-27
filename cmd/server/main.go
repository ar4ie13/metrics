package main

import (
	"log"

	"github.com/ar4ie13/metrics/internal/server/config"
	"github.com/ar4ie13/metrics/internal/server/handler"
	"github.com/ar4ie13/metrics/internal/server/repository"
	"github.com/ar4ie13/metrics/internal/server/service"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	cfg := config.NewConfig()
	repo := repository.NewMemStorage()
	srv := service.NewService(repo)
	hndlr := handler.NewHandler(srv, cfg.HandlerConfig)

	if err := hndlr.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
