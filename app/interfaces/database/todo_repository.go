package database

import "github.com/ymktmk/golang-clean-architecture/app/domain"

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) Store(t *domain.Todo) (todo *domain.Todo, err error) {
	if err = repo.Create(t).Error; err != nil {
		return
	}
	todo = t
	return
}

func (repo *TodoRepository) Update(id int, t *domain.Todo) (todo *domain.Todo, err error) {
	if err = repo.Model(&todo).Where("id = ?", id).Update("name", t.Name).Error; err != nil {
		return
	}
	todo = t
	return
}

func (repo *TodoRepository) FindTodoById(id int) (todo *domain.Todo, err error) {
	if err = repo.Where("id = ?", id).Find(&todo).Error; err != nil {
		return
	}
	return
}

func (repo *TodoRepository) FindTodosById(userId int) (todos *domain.Todos, err error) {
	if err = repo.Where("user_id = ?", userId).Find(&todos).Error; err != nil {
		return
	}
	return
}