package service

import "github.com/XapTMaH19/todo-app/internal/storage"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(storage *storage.Storage) *Service {
	return &Service{}
}
