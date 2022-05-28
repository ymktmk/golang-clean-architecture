package domain

import "time"

type Todos []Todo

type Todo struct {
	ID        uint       `gorm:"primaryKey,autoincrement" json:"id,omitempty"`
	Name      string     `gorm:"size:255;not null" json:"name,omitempty" validate:"required"`
	UserID    uint       `json:"user_id,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" sql:"DEFAULT:current_timestamp on update current_timestamp"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// todo作成
type TodoCreateRequest struct {
	Name string `json:"name" validate:"min=1,max=100"`
}
