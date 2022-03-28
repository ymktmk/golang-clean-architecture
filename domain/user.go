package domain

import "Golang-CleanArchitecture/domain/gorm"

type Users []User

type User struct {
    gorm.Model
    Name  string `gorm:"size:255;not null" json:"name,omitempty" validate:"required"`
    Email string `gorm:"size:255;not null;unique" json:"email,omitempty" validate:"required,email"`
    CompanyID int `gorm:"not null" json:"company_id,omitempty"`
    Todos  []Todo `gorm:"foreignKey:UserID" json:"todos,omitempty"`
}

type UserCreateRequest struct {
    UID  string  `json:"uid" validate:"required"`
    UserName  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    CompanyName string `json:"company" validate:"required"` 
}

type UserUpdateRequest struct {
    UserName  string `json:"name" validate:"required"`
}

type UserUpdateResponse struct {
    UID  string   `json:"uid"`
    UserName  string `json:"name"`
    Email string `json:"email"`
}