package usecase

import (
	"github.com/ymktmk/golang-clean-architecture/app/domain"
)

type UserRepository interface {
	Store(u *domain.User) (user *domain.User, err error)
	Update(uid string, u *domain.User) (user *domain.User, err error)
	// FindByUid(uid string) (user *domain.User, err error)
	FindByEmail(email string) (user *domain.User, err error)
	FindUsersByEmail(email string) (users domain.Users, err error)
	FindById(userId int) (user *domain.User, err error)
	DeleteById(user *domain.User) (err error)
}

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u *domain.User) (user *domain.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	return
}

// func (interactor *UserInteractor) Update(uid string, u *domain.User) (user *domain.User, err error) {
// 	u, err = interactor.UserRepository.Update(uid, u)
// 	user, err = interactor.UserRepository.FindByUid(uid)
// 	return
// }

func (interactor *UserInteractor) UserById(userId int) (user *domain.User, err error) {
	user, err = interactor.UserRepository.FindById(userId)
	return
}

// func (interactor *UserInteractor) UserByUid(uid string) (user *domain.User, err error) {
// 	user, err = interactor.UserRepository.FindByUid(uid)
// 	return
// }

func (interactor *UserInteractor) UserByEmail(email string) (user *domain.User, err error) {
	user, err = interactor.UserRepository.FindByEmail(email)
	return
}

func (interactor *UserInteractor) ExistUserByEmail(email string) (users domain.Users, err error) {
	users, err = interactor.UserRepository.FindUsersByEmail(email)
	return
}
