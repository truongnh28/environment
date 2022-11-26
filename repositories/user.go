package repositories

import (
	"context"
	"github.com/truongnh28/environment-be/dto"
	"github.com/truongnh28/environment-be/models"
	"gorm.io/gorm"
	"time"
)

type UserRepository interface {
	GetAllUsers() ([]models.Users, error)
	GetUserByUsername(ctx context.Context, name string) (models.Users, error)
	Login(ctx context.Context, username string, password string) (models.Users, error)
	Register(ctx context.Context, register dto.User) (models.Users, error)
	ChangeInfo(ctx context.Context, mapData map[string]interface{}) (models.Users, error)
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

func (u *userRepositoryImpl) Login(ctx context.Context, username string, password string) (models.Users, error) {
	var (
		user models.Users
		db   = u.database.WithContext(ctx)
	)
	err := db.Model(models.Users{}).Where("user_name = ? and pass_word = ? ", username, password).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepositoryImpl) Register(ctx context.Context, register dto.User) (models.Users, error) {
	var (
		user models.Users
		db   = u.database.WithContext(ctx)
	)

	user.UserName = register.UserName
	user.PassWord = register.PassWord
	user.IsResolver = register.IsResolver
	user.Phone = register.Phone
	user.Email = register.Email
	_time := time.Now()
	user.CreatedAt = &_time

	err := db.Model(models.Users{}).Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepositoryImpl) ChangeInfo(ctx context.Context, mapData map[string]interface{}) (models.Users, error) {
	var (
		user models.Users
		db   = u.database.WithContext(ctx)
	)
	username := mapData["user_name"].(string)
	if mapData["pass_word"] != nil {
		user.PassWord = mapData["pass_word"].(string)
	}
	mapData["updated_at"] = time.Now()

	err := db.Model(models.Users{}).Where(" user_name = ? ", username).Updates(mapData).Error
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
