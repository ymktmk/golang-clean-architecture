package interactor

import (
	"github.com/ymktmk/golang-clean-architecture/app/domain"
	"github.com/ymktmk/golang-clean-architecture/app/usecase/input"
	"github.com/ymktmk/golang-clean-architecture/app/usecase/output"
)

type TodoRepository interface {
	Store(t *domain.Todo) (todo *domain.Todo, err error)
	Update(id uint, t *domain.Todo) (todo *domain.Todo, err error)
	FindTodoById(id uint) (todo *domain.Todo, err error)
	FindTodosById(userId uint) (todos *domain.Todos, err error)
}

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) Create(inputData input.CreateTodo) (outputData output.Todo, err error) {
	todo := &domain.Todo{
		ID:     inputData.ID,
		Name:   inputData.Name,
		UserID: inputData.UserID,
	}
	newTodo, err := interactor.TodoRepository.Store(todo)
	if err != nil {
		return
	}
	outputData = output.Todo{
		ID:     newTodo.ID,
		Name:   newTodo.Name,
		UserID: newTodo.UserID,
	}
	return
}

func (interactor *TodoInteractor) Update(inputData input.UpdateTodo) (outputData output.Todo, err error) {
	todoId := inputData.ID
	todo := &domain.Todo{
		Name:   inputData.Name,
		UserID: inputData.UserID,
	}
	newTodo, err := interactor.TodoRepository.Update(todoId, todo)
	if err != nil {
		return
	}
	outputData = output.Todo{
		ID:     newTodo.ID,
		Name:   newTodo.Name,
		UserID: newTodo.UserID,
	}
	return
}

func (interactor *TodoInteractor) GetTodo(inputData input.GetTodo) (outputData output.Todo, err error) {
	id := inputData.TodoID
	todo, err := interactor.TodoRepository.FindTodoById(id)
	if err != nil {
		return
	}
	outputData = output.Todo{
		ID:     todo.ID,
		Name:   todo.Name,
		UserID: todo.UserID,
	}
	return
}

func (interactor *TodoInteractor) GetAllTodos(inputData input.GetAllTodos) (outputData output.Todos, err error) {
	userId := inputData.UserID
	todos, err := interactor.TodoRepository.FindTodosById(userId)
	if err != nil {
		return
	}

	outputData = make(output.Todos, 0, len(*todos))

	for _, v := range *todos {
		todo := output.Todo{
			ID:     v.ID,
			Name:   v.Name,
			UserID: v.UserID,
		}
		outputData = append(outputData, todo)
	}
	return
}
