package main

import (
	"github.com/XapTMaH19/todo-app"
	"github.com/XapTMaH19/todo-app/internal/handlers"
	"log"
)

func main() {
	handlers := new(handlers.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8888", handlers.InitRoutes()); err != nil {
		log.Fatalf("eпроизошла ошибка при запуске http-сервера: %s", err.Error())
	}
}
