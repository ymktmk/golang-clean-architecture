package domain

import "github.com/ymktmk/golang-clean-architecture/domain/gorm"

type Users []User

type User struct {
    gorm.Model
    Name  string `gorm:"size:255;not null" json:"name,omitempty" validate:"required"`
    Email string `gorm:"size:255;not null;unique" json:"email,omitempty" validate:"required,email"`
    // Todos  []Todo `gorm:"foreignKey:UserID" json:"todos,omitempty"`
}

type UserCreateRequest struct {
    UserName  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
}

type UserUpdateRequest struct {
    UserName  string `json:"name" validate:"required"`
}

type UserUpdateResponse struct {
    UID  string   `json:"uid"`
    UserName  string `json:"name"`
    Email string `json:"email"`
}