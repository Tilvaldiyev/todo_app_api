package main

import (
	"log"
	"todo_app_api"
	"todo_app_api/pkg/handler"
)

func main() {
	handler := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Error ocured while running http server: %s", err.Error())
	}
}
