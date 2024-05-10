package storage

type Authorization interface {
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

func NewStorage() *Storage {
	return &Storage{}
}
