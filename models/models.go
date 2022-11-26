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
	Author      string     `json:"author" gorm:"column:author"`
	Lag         float32    `json:"lag" gorm:"column:lag"`
	Lng         float32    `json:"lng" gorm:"column:lng"`
	ResolverID  int        `json:"resolver_id" gorm:"column:resolver_id"`
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
	ID         string `json:"id" gorm:"column:id"`
	UserName   string `json:"user_name" gorm:"column:user_name"`
	PassWord   string `json:"pass_word" gorm:"column:pass_word"`
	IsResolver bool   `json:"is_resolver" gorm:"column:is_resolver"`
	Email      string `json:"email" gorm:"column:email"`
	Phone      string `json:"phone" gorm:"column:phone"`
}

// type AccountStatus string

// const (
// 	Active  AccountStatus = "Active"
// 	Blocked AccountStatus = "Blocked"
// )

// type AccountRole string

// const (
// 	SuperAdmin AccountRole = "SuperAdmin"
// 	Admin      AccountRole = "Admin"
// 	User       AccountRole = "User"
// )

// var (
// 	priority = map[AccountRole]int{
// 		SuperAdmin: 0,
// 		Admin:      -1,
// 		User:       -2,
// 	}
// )

// func (a AccountRole) CanChange(ar AccountRole) bool {
// 	return priority[a] > priority[ar]
// }

// func (a AccountRole) IsValid() bool {
// 	_, ok := priority[a]
// 	return ok
// }

// // Status enum
// type Status byte

// const (
// 	InitialStatus = iota
// 	WaitingForScanStatus
// 	ScanningStatus
// 	ScannedStatus
// )

// type Table interface {
// 	TableName() string
// }

// func (Albums) TableName() string {
// 	return "albums"
// }

// func (Artists) TableName() string {
// 	return "artists"
// }

// func (Interactions) TableName() string {
// 	return "interactions"
// }

// func (PlayLists) TableName() string {
// 	return "playlists"
// }

// func (PlayListSongs) TableName() string {
// 	return "playlist_songs"
// }

// func (Songs) TableName() string {
// 	return "songs"
// }

// func (Accounts) TableName() string {
// 	return "accounts"
// }
