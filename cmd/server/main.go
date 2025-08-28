package main

import (
	"github.com/ar4ie13/metrics/internal/config"
	"github.com/ar4ie13/metrics/internal/handler"
	"github.com/ar4ie13/metrics/internal/repository"
	"github.com/ar4ie13/metrics/internal/service"
	"log"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	cfg := config.NewConfig()
	cfg.InitConfig()
	repo := repository.NewMemStorage()
	srv := service.NewService(repo)
	hndlr := handler.NewHandler(srv, cfg)

	if err := hndlr.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
