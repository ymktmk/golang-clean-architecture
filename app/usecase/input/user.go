package input

import (
	"github.com/ymktmk/golang-clean-architecture/app/usecase/output"
)

type UserUsecase interface {
	Create(CreateUser) (output.User, error)
	Update(UpdateUser) (output.User, error)
	UserById(GetUserById) (output.User, error)
	UserByEmail(GetUserByEmail) (output.User, error)
	ExistUserByEmail(GetExistUserByEmail) (output.Users, error)
}

type CreateUser struct {
	ID     uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}

type UpdateUser struct {
	ID     uint   `json:"id"`
	Name     string `json:"name"`
}

type GetUserById struct {
	ID     uint   `json:"id"`
}

type GetUserByEmail struct {
	Email    string `json:"email"`
}

type GetExistUserByEmail struct {
	Email    string `json:"email"`
}