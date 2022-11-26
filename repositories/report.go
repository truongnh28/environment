package repositories

import (
	"context"

	"github.com/truongnh28/environment-be/dto"

	"github.com/truongnh28/environment-be/models"

	ot "github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

type ReportRepository interface {
	Create(ctx context.Context, record *models.Reports) (*models.Reports, error)
	UpdateWithMap(
		octx context.Context,
		record *models.Reports,
		params map[string]interface{},
	) error
	GetByID(octx context.Context, id int) (*models.Reports, error)
	List(octx context.Context, size int, page int, filter *dto.FilterReport) ([]*models.Reports, error)
	Delete(octx context.Context, record *models.Reports) error
	CountWithFilter(octx context.Context, filter *dto.FilterReport) (int64, error)
	GetAll(ctx context.Context) ([]models.Reports, error)
	TopResolver(ctx context.Context) ([]*models.Reports, error)
}

type reportRepositoryImpl struct {
	database *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepositoryImpl{
		database: db,
	}
}

func (r *reportRepositoryImpl) dbWithContext(ctx context.Context) *gorm.DB {
	return r.database.WithContext(ctx)
}

func (r *reportRepositoryImpl) Create(ctx context.Context, record *models.Reports) (*models.Reports, error) {
	err := r.dbWithContext(ctx).Create(record).Error
	return record, err
}

// UpdateWithMap ...
func (r *reportRepositoryImpl) UpdateWithMap(
	octx context.Context,
	record *models.Reports,
	params map[string]interface{},
) error {
	span, ctx := ot.StartSpanFromContext(octx, "ReportSQLRepo_UpdateWithMap")
	defer span.Finish()
	return r.dbWithContext(ctx).
		Model(record).
		Updates(params).
		Error
}

// Delete ...
func (r *reportRepositoryImpl) Delete(octx context.Context, record *models.Reports) error {
	span, ctx := ot.StartSpanFromContext(octx, "ReportSQLRepo_Delete")
	defer span.Finish()

	return r.dbWithContext(ctx).Delete(record).Error
}

func (r *reportRepositoryImpl) buildQueryFromFilter(ctx context.Context, query *gorm.DB, filter *dto.FilterReport) *gorm.DB {
	if filter == nil {
		return query
	}
	if filter.Status != nil {
		query = query.Where("status = ?", *filter.Status)
	}
	if filter.Priority != nil {
		query = query.Where("priority = ?", *filter.Priority)
	}
	if filter.City != nil {
		query = query.Where("city = ?", *filter.City)

	}
	if filter.District != nil {
		query = query.Where("district = ?", *filter.District)
	}
	if filter.Street != nil {
		query = query.Where("street = ?", *filter.Street)
	}
	if filter.Ward != nil {
		query = query.Where("ward = ?", *filter.Ward)
	}
	return query
}

// ListByObject ...
func (r *reportRepositoryImpl) List(
	octx context.Context,
	size int,
	page int,
	filter *dto.FilterReport,
) ([]*models.Reports, error) {
	span, ctx := ot.StartSpanFromContext(octx, "ReportSQLRepo_List")
	defer span.Finish()
	var records []*models.Reports
	query := r.dbWithContext(ctx)
	query = r.buildQueryFromFilter(ctx, query, filter)
	offset := (page - 1) * size
	err := query.Order("created_at DESC").Limit(size).Offset(offset).Find(&records).Error
	return records, err
}

// CountWithFilter
func (r *reportRepositoryImpl) CountWithFilter(octx context.Context, filter *dto.FilterReport) (int64, error) {
	span, ctx := ot.StartSpanFromContext(octx, "ReportSQLRepo_CountWithFilter")
	defer span.Finish()

	var records []*models.Reports
	var c int64

	query := r.dbWithContext(ctx)
	query = r.buildQueryFromFilter(ctx, query, filter)
	err := query.Find(&records).Count(&c).Error
	return c, err
}

func (r *reportRepositoryImpl) GetAll(ctx context.Context) ([]models.Reports, error) {
	var (
		resp = make([]models.Reports, 0)
	)
	err := r.database.WithContext(ctx).Model(models.Reports{}).Find(&resp).Error
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (r *reportRepositoryImpl) GetByID(octx context.Context, id int) (*models.Reports, error) {
	span, ctx := ot.StartSpanFromContext(octx, "ReportSQLRepo_Get")
	defer span.Finish()
	var record *models.Reports

	query := r.dbWithContext(ctx)
	query = query.Where("id = ?", id)
	err := query.Find(&record).Error

	return record, err
}

func (r *reportRepositoryImpl) TopResolver(ctx context.Context) ([]*models.Reports, error) {
	var records []*models.Reports
	query := r.dbWithContext(ctx)
	status := "resolved"
	query = query.Where("status = ?", status)
	err := query.Find(&records).Error

	return records, err
}
