package database

import (
	"github.com/ymktmk/golang-clean-architecture/domain"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u *domain.User) (user *domain.User, err error) {
	if err = repo.Create(u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) Update(uid string, u *domain.User) (user *domain.User, err error) {
	if err = repo.Model(&user).Where("uid = ?", uid).Update("name", u.Name).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) FindByUid(uid string) (user *domain.User, err error) {
	if err = repo.Where("uid = ?", uid).First(&user).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindByEmail(email string) (user *domain.User, err error) {
	if err = repo.Where("email = ?", email).First(&user).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindUsersByEmail(email string) (users domain.Users, err error) {
	if err = repo.Where("email = ?", email).Find(&users).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindById(userId int) (user *domain.User, err error) {
	if err = repo.Where("id = ?", userId).First(&user).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) DeleteById(user *domain.User) (err error) {
	if err = repo.Delete(&user).Error; err != nil {
		return
	}
	return
}
