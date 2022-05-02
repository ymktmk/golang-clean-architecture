package domain

import (
	"github.com/ymktmk/golang-clean-architecture/app/domain/gorm"
)

type Users []User

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null" json:"name,omitempty" validate:"required"`
	Email    string `gorm:"size:255;not null;unique" json:"email,omitempty" validate:"required,email"`
	Password []byte `gorm:"size:255;not null" json:"password,omitempty" validate:"min=8,max=100"`
	// Todos  []Todo `gorm:"foreignKey:UserID" json:"todos,omitempty"`
}

// ユーザー作成
type UserCreateRequest struct {
	UserName string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=8,max=100"`
}

type UserCreateResponse struct {
	gorm.Model
	UserName string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

// ログイン
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=8,max=100"`
}
