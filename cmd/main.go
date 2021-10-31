package main

import (
	"crudEmployee/bootstrap"
	"crudEmployee/config"
	"crudEmployee/internal/handler"
	"crudEmployee/internal/repository"
	"crudEmployee/internal/service"
	"log"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := bootstrap.InitDB(&cfg.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	handle := handler.NewHandler(
		service.NewService(
			repository.NewRepository(db)))

	if err = new(bootstrap.Server).Run(cfg, &handle); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
