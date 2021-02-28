package main

import (
	"github.com/Dimadetected/clientsCRUD"
	"github.com/Dimadetected/clientsCRUD/pkg/handler"
	"github.com/Dimadetected/clientsCRUD/pkg/repository"
	"github.com/Dimadetected/clientsCRUD/pkg/service"
	"log"
)

func main() {

	cfg := repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "mysql",
		DBName:   "postgres",
		SSLMode:  "disabled",
	}

	db, err := repository.NewPostgres(cfg)
	if err != nil {
		log.Print(err)
	}

	rep := repository.NewRepository(db)
	srv := service.NewService(rep)
	handlers := handler.NewHandler(srv)

	server := new(clientsCRUD.Server)
	server.Run("8080", handlers.InitRoutes())
}
