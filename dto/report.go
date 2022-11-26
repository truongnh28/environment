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
	Resolver    string     `json:"resolver,omitempty"`
	City        string     `json:"city,omitempty"`
	District    string     `json:"district,omitempty"`
	Street      string     `json:"street,omitempty"`
	Ward        string     `json:"ward,omitempty"`
	Address     string     `json:"address,omitempty"`
	ImageURL    string     `json:"image_url,omitempty"`
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
	ImageURL    string  `json:"image_url"`
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
	ID       int    `json:"id"`
	Status   string `json:"status"`
	Resolver string `json:"resolver"`
}

type UpdateReportResponse struct {
	StatusError
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
	Title    string  `json:"title"`
	ReportId int     `json:"report_id"`
	Priority int     `json:"priority"`
	Lat      float32 `json:"lat"`
	Lng      float32 `json:"lng"`
}
type MapResponse struct {
	MapResps []MapResp `json:"data"`
	StatusError
}

type GetAllResponse struct {
	Reports []Report
	StatusError
}

type UploadImageResponse struct {
	Url string `json:"url"`
	StatusError
}
