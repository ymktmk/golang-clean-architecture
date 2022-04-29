package usecase

import "github.com/ymktmk/golang-clean-architecture/domain"

type TodoRepository interface {
	FindByUid(uid string) (user *domain.User, err error)
}

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) UserByUid(uid string) (user *domain.User, err error) {
	user, err = interactor.TodoRepository.FindByUid(uid)
	return
}
