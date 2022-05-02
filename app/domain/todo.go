package domain

import "github.com/ymktmk/golang-clean-architecture/app/domain/gorm"

type Todos []Todo

type Todo struct {
	gorm.Model
	Name   string `gorm:"size:255;not null" json:"name,omitempty" validate:"required"`
	UserID int    `json:"user_id,omitempty"`
}

// todo作成
type TodoCreateRequest struct {
	Name string `json:"name" validate:"required"`
}