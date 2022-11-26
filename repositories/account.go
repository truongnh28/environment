package repositories

import (
	"context"
	"spotify/models"

	"gorm.io/gorm"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type ReportRepository interface {
	Create(ctx context.Context, record *models.Reports) (*models.Reports, error)
	UpdateWithMap(ctx context.Context, record *models.Reports, params map[string]interface{}) error
	GetByID(ctx context.Context, id int) (*models.Reports, error)
	List(ctx context.Context, page int, size int, filter *models.ReportFilter) ([]*models.Reports, error)
}

type accountRepositoryImpl struct {
	database *gorm.DB
}

func (a *accountRepositoryImpl) GetByDomain(name string) ([]*models.Accounts, error) {
	userProfiles := make([]*models.Accounts, 0)
	database := a.database

	database = database.Model(models.Accounts{})

	if name != "" {
		database = database.Where("user_name like ?", "%"+name+"%")
	}
	err := database.Find(&userProfiles).Error
	if err != nil {
		return userProfiles, err
	}
	return userProfiles, nil
}

func (a *accountRepositoryImpl) Create(ctx context.Context, acc models.Accounts) error {
	var (
		db = a.database.WithContext(ctx)
	)
	return db.Model(models.Accounts{}).Create(&acc).Error
}

func (a *accountRepositoryImpl) FindByUserName(ctx context.Context, username string) (models.Accounts, error) {
	var (
		users models.Accounts
		db    = a.database.WithContext(ctx)
	)
	err := db.Model(models.Accounts{}).First(&users).Error
	if err != nil {
		return models.Accounts{}, err
	}
	return users, nil
}

func (a *accountRepositoryImpl) UpdateByUsername(ctx context.Context, username string, acc models.Accounts) (int64, error) {
	db := a.database.WithContext(ctx)
	result := db.Model(models.Accounts{}).Select("status", "role").Where("user_name = ?", username).Updates(acc)
	return result.RowsAffected, result.Error
}

func NewAccountRepository(database *gorm.DB) AccountRepository {
	return &accountRepositoryImpl{
		database: database,
	}
}
