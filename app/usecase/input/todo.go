package input

import (
	"github.com/ymktmk/golang-clean-architecture/app/usecase/output"
)

type TodoUserCase interface {
	Craete(CreateTodo) (output.CreateTodo, error)
	Update(UpdateTodo) (output.UpdateTodo, error)
	GetTodo(GetTodo) (output.GetTodo, error)
	GetAllTodos(GetAllTodos) (output.GetAllTodos, error)
}

type CreateTodo struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
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
