package main

import (
	"log"
	"todo_app_api"
	"todo_app_api/pkg/handler"
	"todo_app_api/pkg/repository"
	"todo_app_api/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error ocured while running http server: %s", err.Error())
	}
}
