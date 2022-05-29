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
	UserID uint   `json:"user_id"`
}

type UpdateTodo struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
}

type GetTodo struct {
	TodoID uint `json:"todo_id"`
}

type GetAllTodos struct {
	UserID uint `json:"user_id"`
}
