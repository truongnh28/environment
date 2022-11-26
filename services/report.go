package services

import (
	"context"
	"github.com/golang/glog"
	"github.com/truongnh28/environment-be/dto"
	"github.com/truongnh28/environment-be/helper/common"
	"github.com/truongnh28/environment-be/repositories"
)

//Create(ctx context.Context, record *models.Reports) (*models.Reports, error)
//UpdateWithMap(
//	octx context.Context,
//	record *models.Reports,
//	params map[string]interface{},
//) error
//GetByID(ctx context.Context, id int) (*models.Reports, error)
//List(
//	octx context.Context,
//	size int,
//	page int,
//	filter *models.ReportFilter,
//) ([]*models.Reports, error)
//Delete(octx context.Context, record *models.Reports) error
//CountWithFilter(
//	octx context.Context,
//	filter *models.ReportFilter,
//) (int64, error)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type ReportService interface {
	Create(ctx context.Context, message *dto.CreateReportRequest) (*dto.CreateReportResponse, error)
	GetByID(ctx context.Context, message *dto.GetReportByIDRequest) (*dto.GetReportByIDResponse, error)
	List(ctx context.Context, message *dto.ListReportsRequest) (*dto.ListReportsResponse, error)
	// UpdateWithMap(ctx context.Context, *dto.UpdateReportRequest) error
}

func NewReportService(reportRepository repositories.ReportRepository) ReportService {
	return &reportServiceImpl{
		reportRepository: reportRepository,
	}
}

type reportServiceImpl struct {
	reportRepository repositories.ReportRepository
}



func (r *reportService) Create(ctx context.Context, message *dto.CreateReportRequest) (*dto.CreateReportResponse, error) {
	record := converter.FromReportDTO(message)
	report, err := r.reportRepo.Create(ctx, record)
	if err != nil {
		return nil, err
	}

	resp := &dto.CreateReportResponse{
		Report: *converter.ToReportDTO(report),
	}
	return resp, nil
}

func (r *reportService) List(ctx context.Context, message *dto.ListReportsRequest) (*dto.ListReportsResponse, error) {

}

// func (r *reportService) UpdateWithMap(ctx context.Context, message *dto.UpdateReportRequest) error {
// 	params := map[string]interface{}
// 	for v := range message.FieldMask {
// 		params[v] =
// 	}
// 	err := r.reportRepo.UpdateWithMap(octx, record.)
// }

func (r *reportServiceImpl) Create(ctx context.Context, message *dto.CreateReportRequest) (*dto.CreateReportResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *reportServiceImpl) GetByID(ctx context.Context, message dto.GetReportByIDRequest) dto.GetReportByIDResponse {
	//TODO implement me
	panic("implement me")
}

func (r *reportServiceImpl) Update(ctx context.Context, username string, req dto.UpdateAccountRequest) common.SubReturnCode {
	//TODO implement me
	panic("implement me")
}

func (r *reportServiceImpl) List(ctx context.Context, message *dto.ListReportsRequest) (*dto.ListReportsResponse, error) {
	var (
		resp    dto.ListReportsResponse
		reports = make([]*dto.Report, 0)
	)
	reportResp, err := r.reportRepository.List(ctx, message.Size, message.Page, message.Filter)
	if err != nil {
		glog.Errorf("Get List in report service err: ", err)
	}

	for _, val := range reportResp {
		reports = append(reports, &dto.Report{
			ID:          val.ID,
			Title:       val.Title,
			Description: val.Description,
			CreatedAt:   val.CreatedAt,
			UpdatedAt:   val.UpdatedAt,
			DeletedAt:   val.DeletedAt,
			Status:      val.Status,
			Priority:    val.Priority,
			Author:      val.Author,
			Lag:         val.Lag,
			Lng:         val.Lng,
			ResolverID:  val.ResolverID,
			City:        val.City,
			District:    val.District,
			Street:      val.Street,
			Ward:        val.Ward,
			Address:     val.Address,
		})
	}
	resp.Reports = reports
	resp.Size = message.Size
	resp.Page = message.Page
	return &resp, nil
}

