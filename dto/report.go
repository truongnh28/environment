package dto

import "time"

type Report struct {
	ID          int        `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Status      string     `json:"status,omitempty"`
	Priority    string     `json:"priority,omitempty"`
	Author      string     `json:"author,omitempty"`
	Lag         float32    `json:"lag,omitempty"`
	Lng         float32    `json:"lng,omitempty"`
	ResolverID  int        `json:"resolver_id,omitempty"`
	City        string     `json:"city,omitempty"`
	District    string     `json:"district,omitempty"`
	Street      string     `json:"street,omitempty"`
	Ward        string     `json:"ward,omitempty"`
	Address     string     `json:"address,omitempty"`
}

type CreateReportRequest struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Priority    string  `json:"priority,omitempty"`
	UserName    string  `json:"user_name,omitempty"`
	City        string  `json:"city,omitempty"`
	District    string  `json:"district,omitempty"`
	Street      string  `json:"street,omitempty"`
	Ward        string  `json:"ward,omitempty"`
	Address     *string `json:"address,omitempty"`
}

type CreateReportResponse struct {
	report Report
}

type ListReportsRequest struct {
	Page   int           `json:"page,omitempty"`
	Size   int           `json:"size,omitempty"`
	Filter *FilterReport `json:"filter,omitempty"`
}

type ListReportsResponse struct {
	Page       int       `json:"page,omitempty"`
	Size       int       `json:"size,omitempty"`
	Reports    []*Report `json:"reports,omitempty"`
	TotalItems int       `json:"total_items,omitempty"`
}

type FilterReport struct {
	Status   *string `json:"status"`
	Priority *string `json:"priority"`
	City     *string `json:"city"`
	District *string `json:"district"`
	Street   *string `json:"street"`
	Ward     *string `json:"ward"`
}

type UpdateReportRequest struct {
	Record    Report   `json:"record"`
	FieldMask []string `json:"FieldMask"`
}

type UpdateReportResponse struct {
	report Report
}

type GetReportByIDRequest struct {
	ID int `json:"id,omitempty"`
}

type GetReportByIDResponse struct {
	Report      Report     `json:"report,omitempty"`
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

type MapResp struct {
	ReportId int     `json:"report_id"`
	Priority string  `json:"priority"`
	Lat      float32 `json:"lat"`
	Lng      float32 `json:"lng"`
}
type MapResponse struct {
	MapResps []MapResp `json:"data"`
	StatusError
}
