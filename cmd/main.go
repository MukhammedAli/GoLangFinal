package main

import (
	"log"

	todo "github.com/MukhammedAli/GoFinalProject"
	"github.com/MukhammedAli/GoFinalProject/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while http server: %s", err.Error())
	}
}
