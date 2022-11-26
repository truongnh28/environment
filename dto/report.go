package dto

import "time"

type Report struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Status      string     `json:"status"`
	Author      string     `json:"author"`
	Lag         float32    `json:"lag"`
	Lng         float32    `json:"lng"`
	ResolverID  int        `json:"resolver_id"`
}

type ReportResponse struct {
	Reports []Report
	StatusError
}
