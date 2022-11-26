package services

import (
	"context"
	"sort"

	"github.com/truongnh28/environment-be/helper/common"

	"github.com/golang/glog"
	"github.com/truongnh28/environment-be/dto"
	"github.com/truongnh28/environment-be/pkg/converter"
	"github.com/truongnh28/environment-be/repositories"
)

type ReportService interface {
	Create(ctx context.Context, message dto.CreateReportRequest) (dto.CreateReportResponse, common.SubReturnCode)
	GetByID(ctx context.Context, message dto.GetReportByIDRequest) (dto.GetReportByIDResponse, common.SubReturnCode)
	List(ctx context.Context, message dto.ListReportsRequest) (dto.ListReportsResponse, common.SubReturnCode)
	MapReportList(ctx context.Context) ([]dto.MapResp, common.SubReturnCode)
	Update(ctx context.Context, message dto.UpdateReportRequest) common.SubReturnCode
	GetAll(ctx context.Context) (dto.GetAllResponse, common.SubReturnCode)
	TopResolver(ctx context.Context) (dto.TopResolverResponse, common.SubReturnCode)
}

func NewReportService(reportRepository repositories.ReportRepository) ReportService {
	return &reportServiceImpl{
		reportRepository: reportRepository,
	}
}

type reportServiceImpl struct {
	reportRepository repositories.ReportRepository
}

func (r *reportServiceImpl) Create(ctx context.Context, message dto.CreateReportRequest) (dto.CreateReportResponse, common.SubReturnCode) {
	record := converter.FromReportDTO(&message)
	record.Status = "draft"
	record.Resolver = ""
	report, err := r.reportRepository.Create(ctx, record)
	if err != nil {
		glog.Infoln("Create service err: ", err)
		return dto.CreateReportResponse{}, common.SystemError
	}

	resp := dto.CreateReportResponse{
		Report: *converter.ToReportDTO(report),
	}
	return resp, common.OK
}

func (r *reportServiceImpl) Update(ctx context.Context, message dto.UpdateReportRequest) common.SubReturnCode {
	record, err := r.reportRepository.GetByID(ctx, message.ID)
	if err != nil {
		return common.SystemError
	}
	params := map[string]interface{}{}
	params["status"] = message.Status
	params["resolver_id"] = message.Resolver
	err = r.reportRepository.UpdateWithMap(ctx, record, params)
	if err != nil {
		return common.SystemError
	}
	return common.OK
}

func (r *reportServiceImpl) GetByID(ctx context.Context, message dto.GetReportByIDRequest) (dto.GetReportByIDResponse, common.SubReturnCode) {
	record, err := r.reportRepository.GetByID(ctx, message.ID)
	if err != nil {
		glog.Infoln("Create service err: ", err)
		return dto.GetReportByIDResponse{}, common.SystemError
	}

	return dto.GetReportByIDResponse{
		Report: *converter.ToReportDTO(record),
	}, common.OK
}

func (r *reportServiceImpl) List(ctx context.Context, message dto.ListReportsRequest) (dto.ListReportsResponse, common.SubReturnCode) {
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
			Resolver:    val.Resolver,
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
	countResp, err := r.reportRepository.CountWithFilter(ctx, message.Filter)
	if err != nil {
		glog.Errorf("Get List in report service err: ", err)
	}
	resp.TotalItems = int(countResp)
	return resp, common.OK
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
			Title:    val.Title,
			ReportId: val.ID,
			Priority: val.Priority,
			Lat:      val.Lag,
			Lng:      val.Lng,
		})
	}
	return resp, common.OK
}

func (r *reportServiceImpl) GetAll(ctx context.Context) (dto.GetAllResponse, common.SubReturnCode) {
	resp, err := r.reportRepository.GetAll(ctx)
	if err != nil {
		return dto.GetAllResponse{}, common.SystemError
	}
	var reports []dto.Report
	for _, val := range resp {
		reports = append(reports, *converter.ToReportDTO(&val))
	}
	return dto.GetAllResponse{
		Reports: reports,
	}, common.OK
}

func (r *reportServiceImpl) TopResolver(ctx context.Context) (dto.TopResolverResponse, common.SubReturnCode) {
	records, err := r.reportRepository.TopResolver(ctx)
	if err != nil {
		return dto.TopResolverResponse{}, common.SystemError
	}
	topResolver := map[string]int{}
	for _, v := range records {
		if _, ok := topResolver[v.Resolver]; !ok {
			topResolver[v.Resolver] = 1
		} else {
			topResolver[v.Resolver] += 1
		}
	}
	keys := make([]string, 0, len(topResolver))

	for key := range topResolver {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return topResolver[keys[i]] > topResolver[keys[j]]
	})

	userNames := []string{}
	totals := []int{}
	dem := 0
	for _, k := range keys {
		dem += 1
		if dem == 10 {
			break
		}
		userNames = append(userNames, k)
		totals = append(totals, topResolver[k])
	}
	return dto.TopResolverResponse{
		UserNames: userNames,
		Totals:    totals,
	}, common.OK
}
