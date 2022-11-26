package models

import (
	"time"

	"gorm.io/gorm"
)

type Reports struct {
	*gorm.Model
	ID          int        `json:"id" gorm:"column:id"`
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	Status      string     `json:"status" gorm:"column:status"`
	Priority    int        `json:"priority" gorm:"column:priority"`
	Author      string     `json:"author" gorm:"column:author"`
	Lag         float32    `json:"lag" gorm:"column:lag"`
	Lng         float32    `json:"lng" gorm:"column:lng"`
	Resolver    string     `json:"resolver" gorm:"column:resolver"`
	City        string     `json:"city" gorm:"column:city"`
	District    string     `json:"district" gorm:"column:district"`
	Street      string     `json:"street" gorm:"column:street"`
	Ward        string     `json:"ward" gorm:"column:ward"`
	Address     string     `json:"address" gorm:"column:address"`
	ImageURL    string     `json:"image_url" gorm:"column:image_url"`
}

type Images struct {
	*gorm.Model
	ID       string `json:"id" gorm:"column:id"`
	ReportID int    `json:"report_id" gorm:"column:report_id"`
	Url      string `json:"url" gorm:"column:url"`
	Status   int    `json:"status" gorm:"column:status"`
}

type Users struct {
	*gorm.Model
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

type Table interface {
	TableName() string
}

func (Users) TableName() string {
	return "users"
}

func (Images) TableName() string {
	return "images"
}

func (Reports) TableName() string {
	return "reports"
}
