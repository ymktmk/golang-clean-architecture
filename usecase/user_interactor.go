// ロジックをここに書く
package usecase

import (
	"Golang-CleanArchitecture/domain"
	"fmt"
)

type UserInteractor struct {
    UserRepository UserRepository
}

func (interactor *UserInteractor) Add(user domain.User) (domain.User, error) {
    identifier, err := interactor.UserRepository.Store(user)
    if err != nil {
        fmt.Println(err)
    }
    user, err = interactor.UserRepository.FindById(identifier)
    return user, err
}