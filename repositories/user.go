package repositories

import (
	"context"
	"gorm.io/gorm"
	"spotify/models"
)

type UserRepository interface {
	GetAllUsers() ([]models.Users, error)
	GetUserByUsername(ctx context.Context, name string) (models.Users, error)
}

type userRepositoryImpl struct {
	database *gorm.DB
}

func (u *userRepositoryImpl) GetAllUsers() ([]models.Users, error) {
	resp := make([]models.Users, 0)
	err := u.database.Model(models.Users{}).Find(&resp).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (u *userRepositoryImpl) GetUserByUsername(ctx context.Context, name string) (models.Users, error) {
	var (
		user models.Users
		db   = u.database.WithContext(ctx)
	)
	err := db.Model(models.Users{}).Where("user_name like ?", "%"+name+"%").First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		database: database,
	}
}
