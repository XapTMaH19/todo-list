package storage

import (
	"github.com/XapTMaH19/todo-app/internal/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username string, password string) (models.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Storage struct {
	Authorization
	TodoList
	TodoItem
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorization: NewAuthPostgres(db),
	}
}
