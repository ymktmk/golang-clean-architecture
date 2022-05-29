package interactor

import (
	"github.com/ymktmk/golang-clean-architecture/app/domain"
)

type UserRepository interface {
	Store(u *domain.User) (user *domain.User, err error)
	Update(id int, u *domain.User) (user *domain.User, err error)
	FindByEmail(email string) (user *domain.User, err error)
	FindUsersByEmail(email string) (users domain.Users, err error)
	FindById(userId int) (user *domain.User, err error)
	DeleteById(user *domain.User) (err error)
}

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Create(u *domain.User) (user *domain.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Update(id int, u *domain.User) (user *domain.User, err error) {
	user, err = interactor.UserRepository.Update(id, u)
	return
}

func (interactor *UserInteractor) UserById(userId int) (user *domain.User, err error) {
	user, err = interactor.UserRepository.FindById(userId)
	return
}

func (interactor *UserInteractor) UserByEmail(email string) (user *domain.User, err error) {
	user, err = interactor.UserRepository.FindByEmail(email)
	return
}

func (interactor *UserInteractor) ExistUserByEmail(email string) (users domain.Users, err error) {
	users, err = interactor.UserRepository.FindUsersByEmail(email)
	return
}
