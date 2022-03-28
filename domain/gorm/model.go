package gorm

import (
	"time"
)

type Model struct {
	ID        uint        `gorm:"primaryKey,autoincrement" json:"id,omitempty"`
	CreatedAt *time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time  `json:"deleted_at,omitempty"`
}