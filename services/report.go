package services

import (
	"context"
	"github.com/truongnh28/environment-be/helper/common"

	"github.com/golang/glog"
	"github.com/truongnh28/environment-be/dto"
	"github.com/truongnh28/environment-be/pkg/converter"
	"github.com/truongnh28/environment-be/repositories"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type ReportService interface {
	Create(ctx context.Context, message *dto.CreateReportRequest) (*dto.CreateReportResponse, error)
	GetByID(ctx context.Context, message *dto.GetReportByIDRequest) (*dto.GetReportByIDResponse, error)
	List(ctx context.Context, message *dto.ListReportsRequest) (*dto.ListReportsResponse, error)
	MapReportList(ctx context.Context) ([]dto.MapResp, common.SubReturnCode)
	Update(ctx context.Context, message *dto.UpdateReportRequest) error
}

func NewReportService(reportRepository repositories.ReportRepository) ReportService {
	return &reportServiceImpl{
		reportRepository: reportRepository,
	}
}

type reportServiceImpl struct {
	reportRepository repositories.ReportRepository
}

func (r *reportServiceImpl) Create(ctx context.Context, message *dto.CreateReportRequest) (*dto.CreateReportResponse, error) {
	record := converter.FromReportDTO(message)
	report, err := r.reportRepository.Create(ctx, record)
	if err != nil {
		return nil, err
	}

	resp := &dto.CreateReportResponse{
		Report: *converter.ToReportDTO(report),
	}
	return resp, nil
}

func (r *reportServiceImpl) Update(ctx context.Context, message *dto.UpdateReportRequest) error {
	record := converter.FromReport(message.Record)
	var params map[string]interface{}
	params["status"] = message.Status
	params["resolver_id"] = message.Resolver
	err := r.reportRepository.UpdateWithMap(ctx, record, params)

	return err
}

func (r *reportServiceImpl) GetByID(ctx context.Context, message *dto.GetReportByIDRequest) (*dto.GetReportByIDResponse, error) {
	record, err := r.reportRepository.GetByID(ctx, message.ID)
	if err != nil {
		return nil, err
	}

	return &dto.GetReportByIDResponse{
		Report: *converter.ToReportDTO(record),
	}, nil
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

func (r *reportServiceImpl) MapReportList(ctx context.Context) ([]dto.MapResp, common.SubReturnCode) {
	var (
		resp = make([]dto.MapResp, 0)
	)
	respReports, err := r.reportRepository.GetAll(ctx)
	if err != nil {
		glog.Errorf("MapResponseList err: ", err)
		return nil, common.SystemError
	}
	for _, val := range respReports {
		resp = append(resp, dto.MapResp{
			ReportId: val.ID,
			Priority: val.Priority,
			Lat:      val.Lag,
			Lng:      val.Lng,
		})
	}
	return resp, common.OK
}
