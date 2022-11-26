package services

import (
	"context"
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
	Update(ctx context.Context, username string, req dto.UpdateAccountRequest) common.SubReturnCode
	Create(ctx context.Context, message dto.CreateReportRequest) dto.CreateReportResponse
	FindByUserName(ctx context.Context, username string) (dto.Account, common.SubReturnCode)
}

func NewReportService(reportRepository repositories.ReportRepository) ReportService {
	return &reportServiceImpl{
		reportRepository: reportRepository,
	}
}

type reportServiceImpl struct {
	reportRepository repositories.ReportRepository
}

func (r reportServiceImpl) Update(ctx context.Context, username string, req dto.UpdateAccountRequest) common.SubReturnCode {
	//TODO implement me
	panic("implement me")
}

func (r reportServiceImpl) Create(ctx context.Context, message dto.CreateReportRequest) dto.CreateReportResponse {
	//TODO implement me
	panic("implement me")
}

func (r reportServiceImpl) FindByUserName(ctx context.Context, username string) (dto.Account, common.SubReturnCode) {
	//TODO implement me
	panic("implement me")
}
