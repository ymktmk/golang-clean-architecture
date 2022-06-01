package interactor

import (
	"github.com/ymktmk/golang-clean-architecture/app/domain"
	"github.com/ymktmk/golang-clean-architecture/app/usecase/input"
	"github.com/ymktmk/golang-clean-architecture/app/usecase/output"
)

type UserRepository interface {
	Store(u *domain.User) (user *domain.User, err error)
	Update(id int, u *domain.User) (user *domain.User, err error)
	FindByEmail(email string) (user *domain.User, err error)
	FindUsersByEmail(email string) (users *domain.Users, err error)
	FindById(userId int) (user *domain.User, err error)
	DeleteById(user *domain.User) (err error)
}

type UserInteractor struct {
	UserRepository UserRepository
}

func NewUserInteractor(userRepo UserRepository) *UserInteractor {
	return &UserInteractor{
		UserRepository: userRepo,
	}
}

func (interactor *UserInteractor) Create(inputData input.CreateUser) (outputData output.User, err error) {
	user := &domain.User{
		Name: inputData.Name,
		Email: inputData.Email,
		Password: inputData.Password,
	}
	newUser, err := interactor.UserRepository.Store(user)
	if err != nil {
		return
	}
	outputData = output.User{
		ID: newUser.ID,
		Name: newUser.Name,
		Email: newUser.Email,
	}
	return
}

func (interactor *UserInteractor) Update(inputData input.UpdateUser) (outputData output.User, err error) {
	id := int(inputData.ID)
	user := &domain.User{
		Name: inputData.Name,
	}
	newUser, err := interactor.UserRepository.Update(id, user)
	if err != nil {
		return
	}
	outputData = output.User{
		ID: newUser.ID,
		Name: newUser.Name,
		Email: newUser.Email,
	}
	return
}

func (interactor *UserInteractor) UserById(inputData input.GetUserById) (outputData output.User, err error) {
	id := int(inputData.ID)
	newUser, err := interactor.UserRepository.FindById(id)
	if err != nil {
		return
	}
	outputData = output.User{
		ID: newUser.ID,
		Name: newUser.Name,
		Email: newUser.Email,
	}
	return
}

func (interactor *UserInteractor) UserByEmail(inputData input.GetUserByEmail) (outputData output.User, err error) {
	email := inputData.Email
	user, err := interactor.UserRepository.FindByEmail(email)
	if err != nil {
		return
	}
	outputData = output.User{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}
	return
}

func (interactor *UserInteractor) ExistUserByEmail(inputData input.GetExistUserByEmail) (outputData output.Users, err error) {
	email := inputData.Email
	users, err := interactor.UserRepository.FindUsersByEmail(email)
	if err != nil {
		return
	}

	outputData = make(output.Users, 0, len(*users))

	for _, v := range *users {
		user := output.User{
			ID:     v.ID,
			Name:   v.Name,
			Email: v.Email,
		}
		outputData = append(outputData, user)
	}
	return
}
