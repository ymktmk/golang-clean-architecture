// ロジックをここに書く
package usecase

import (
	"Golang-CleanArchitecture/domain"
)

type UserInteractor struct {
    UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (domain.User, error) {
    identifier, err := interactor.UserRepository.Store(u)
    user, err := interactor.UserRepository.FindById(identifier)
    return user, err
}

func (interactor *UserInteractor) Users() (domain.Users, error) {
    users, err := interactor.UserRepository.FindAll()
    return users, err
}

func (interactor *UserInteractor) Show(id int) (domain.User, error) {
    user, err := interactor.UserRepository.FindById(id)
    return user, err
}


