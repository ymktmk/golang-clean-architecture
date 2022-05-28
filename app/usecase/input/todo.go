package input

import (
	"github.com/ymktmk/golang-clean-architecture/app/usecase/output"
)

type TodoUsecase interface {
	Create(CreateTodo) (output.Todo, error)
	Update(UpdateTodo) (output.Todo, error)
	GetTodo(GetTodo) (output.Todo, error)
	GetAllTodos(GetAllTodos) (output.Todo, error)
}

type CreateTodo struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}

type UpdateTodo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

type GetTodo struct {
	TodoID int `json:"todo_id"`
}

type GetAllTodos struct {
	UserID int `json:"user_id"`
}
