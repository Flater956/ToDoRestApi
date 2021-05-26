package app

import (
	"ToDoRestApi/internal/server"
	"ToDoRestApi/internal/transfer/http"
	"log"
)

func Run() {
	server := new(server.Server)
	handlers := new(http.Handler)

	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("server error: %s", err.Error())
	}
}
