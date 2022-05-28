package interactor

import "github.com/ymktmk/golang-clean-architecture/app/domain"

type TodoRepository interface {
	Store(t *domain.Todo) (todo *domain.Todo, err error)
	Update(id int, t *domain.Todo) (todo *domain.Todo, err error)
	FindTodoById(id int) (todo *domain.Todo, err error)
	FindTodosById(userId int) (todos *domain.Todos, err error)
}

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) Create(t *domain.Todo) (todo *domain.Todo, err error) {
	todo, err = interactor.TodoRepository.Store(t)
	return
}

func (interactor *TodoInteractor) Update(id int, t *domain.Todo) (todo *domain.Todo, err error) {
	todo, err = interactor.TodoRepository.Update(id, t)
	return
}

func (interactor *TodoInteractor) GetTodo(id int) (todo *domain.Todo, err error) {
	todo, err = interactor.TodoRepository.FindTodoById(id)
	return
}

func (interactor *TodoInteractor) GetAllTodos(userId int) (todos *domain.Todos, err error) {
	todos, err = interactor.TodoRepository.FindTodosById(userId)
	return
}
