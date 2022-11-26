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
	Priority    int        `json:"priority,omitempty"`
	Author      string     `json:"author,omitempty"`
	Lag         float32    `json:"lag,omitempty"`
	Lng         float32    `json:"lng,omitempty"`
	ResolverID  int        `json:"resolver_id,omitempty"`
	City        string     `json:"city,omitempty"`
	District    string     `json:"district,omitempty"`
	Street      string     `json:"street,omitempty"`
	Ward        string     `json:"ward,omitempty"`
	Address     string     `json:"address,omitempty"`
	Images      []string   `json:"images,omitempty"`
}

type CreateReportRequest struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Priority    int     `json:"priority,omitempty"`
	UserName    string  `json:"user_name,omitempty"`
	City        string  `json:"city,omitempty"`
	District    string  `json:"district,omitempty"`
	Street      string  `json:"street,omitempty"`
	Ward        string  `json:"ward,omitempty"`
	Address     *string `json:"address,omitempty"`
	Lag         float32 `json:"lag,omitemty"`
	Lng         float32 `json:"lng,omitempty"`
}

type CreateReportResponse struct {
	Report Report `json:"report"`
	StatusError
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
	Priority *int    `json:"priority"`
	City     *string `json:"city"`
	District *string `json:"district"`
	Street   *string `json:"street"`
	Ward     *string `json:"ward"`
}

type UpdateReportRequest struct {
	Record   Report `json:"record"`
	Status   string `json:"status"`
	Resolver int    `json:"resolver"`
}

type UpdateReportResponse struct {
	report Report
}

type GetReportByIDRequest struct {
	ID int `json:"id,omitempty"`
}

type GetReportByIDResponse struct {
	Report Report `json:"report,omitempty"`
	StatusError
}

type ReportResponse struct {
	Reports []Report
	StatusError
}

type MapResp struct {
	ReportId int     `json:"report_id"`
	Priority int     `json:"priority"`
	Lat      float32 `json:"lat"`
	Lng      float32 `json:"lng"`
}
type MapResponse struct {
	MapResps []MapResp `json:"data"`
	StatusError
}
