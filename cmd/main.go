package main

import (
	"fmt"
	"github.com/Dimadetected/clientsCRUD"
	"github.com/Dimadetected/clientsCRUD/pkg/handler"
	"github.com/Dimadetected/clientsCRUD/pkg/repository"
	"github.com/Dimadetected/clientsCRUD/pkg/service"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	cfg := repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "mysql",
		DBName:   "postgres",
		SSLMode:  "disable",
	}

	db, err := repository.NewPostgres(cfg)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(db)
	rep := repository.NewRepository(db)
	srv := service.NewService(rep)
	handlers := handler.NewHandler(srv)

	server := new(clientsCRUD.Server)
	server.Run("8080", handlers.InitRoutes())
}
