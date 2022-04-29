package database

import "github.com/ymktmk/golang-clean-architecture/app/domain"

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) FindByUid(uid string) (user *domain.User, err error) {
	if err = repo.Where("uid=?", uid).First(&user).Error; err != nil {
		return
	}
	return
}
