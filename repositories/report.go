package repositories

import (
	"context"
	"spotify/models"

	ot "github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type ReportRepository interface {
	Create(ctx context.Context, record *models.Reports) (*models.Reports, error)
	UpdateWithMap(
		octx context.Context,
		record *models.Reports,
		params map[string]interface{},
	) error
	GetByID(ctx context.Context, id int) (*models.Reports, error)
	List(
		octx context.Context,
		size int,
		page int,
		filter *models.ReportFilter,
	) ([]*models.Reports, error)
	Delete(octx context.Context, record *models.Reports) error
	CountWithFilter(
		octx context.Context,
		filter *models.ReportFilter,
	) (int64, error)
}

type ReportSQLRepo struct {
	db *gorm.DB
}

func NewReportSQLRepo(db *gorm.DB) *ReportSQLRepo {
	return &ReportSQLRepo{
		db: db,
	}
}

func (r *ReportSQLRepo) dbWithContext(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx)
}

func (r *ReportSQLRepo) Create(ctx context.Context, record *models.Reports) (*models.Reports, error) {
	err := r.dbWithContext(ctx).Create(record).Error
	return record, err
}

// UpdateWithMap ...
func (r *ReportSQLRepo) UpdateWithMap(
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
func (r *ReportSQLRepo) Delete(octx context.Context, record *models.Reports) error {
	span, ctx := ot.StartSpanFromContext(octx, "ReportSQLRepo_Delete")
	defer span.Finish()

	return r.dbWithContext(ctx).Delete(record).Error
}

func (r *ReportSQLRepo) buildQueryFromFilter(
	ctx context.Context,
	query *gorm.DB,
	filter *models.ReportFilter,
) *gorm.DB {
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
func (r *ReportSQLRepo) List(
	octx context.Context,
	size int,
	page int,
	filter *models.ReportFilter,
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
func (r *ReportSQLRepo) CountWithFilter(
	octx context.Context,
	filter *models.ReportFilter,
) (int64, error) {
	span, ctx := ot.StartSpanFromContext(octx, "ReportSQLRepo_CountWithFilter")
	defer span.Finish()

	var records []*models.Reports
	var c int64

	query := r.dbWithContext(ctx)
	query = r.buildQueryFromFilter(ctx, query, filter)
	err := query.Find(&records).Count(&c).Error
	return c, err
}
