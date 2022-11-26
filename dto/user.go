package dto

import (
	"time"
)

type User struct {
	ID         string     `json:"id" gorm:"column:id"`
	UserName   string     `json:"user_name" gorm:"column:user_name"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	PassWord   string     `json:"pass_word" gorm:"column:pass_word"`
	IsResolver bool       `json:"is_resolver" gorm:"column:is_resolver"`
	Email      string     `json:"email" gorm:"column:email"`
	Phone      string     `json:"phone" gorm:"column:phone"`
}

type UserResponse struct {
	Users []User
	StatusError
}

type Register struct {
	UserName   string `json:"username" gorm:"column:user_name"`
	Password   string `json:"password" gorm:"column:pass_word"`
	IsResolver bool   `json:"is_resolver" gorm:"column:is_resolver"`
	Email      string `json:"email,omitempty" gorm:"column:email"`
	Phone      string `json:"phone,omitempty" gorm:"column:phone"`
}
