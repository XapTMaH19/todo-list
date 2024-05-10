package main

import (
	"github.com/XapTMaH19/todo-app"
	"github.com/XapTMaH19/todo-app/internal/handler"
	"github.com/XapTMaH19/todo-app/internal/service"
	"github.com/XapTMaH19/todo-app/internal/storage"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %v", err.Error())
	}
	storage := storage.NewStorage()
	services := service.NewService(storage)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("eпроизошла ошибка при запуске http-сервера: %v", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
