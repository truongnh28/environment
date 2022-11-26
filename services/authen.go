package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/glog"
	"github.com/truongnh28/environment-be/config"
	"github.com/truongnh28/environment-be/dto"
	"github.com/truongnh28/environment-be/repositories"
	"github.com/truongnh28/environment-be/utils/auth"
)

type AuthenService interface {
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponseData, error)
}

func (a *authenServiceImpl) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponseData, error) {
	return a.checkUserFromDB(ctx, req)
}

func (a *authenServiceImpl) checkUserFromDB(ctx context.Context, userInfo dto.LoginRequest) (*dto.LoginResponseData, error) {
	var (
		resp    *dto.LoginResponseData
		account *dto.Account
	)

	acc, err := a.userRepo.GetUserByUsername(ctx, userInfo.Username)
	if err != nil {
		glog.Errorln("GetUserErr: ", err)
		return nil, err
	}
	glog.Infoln("AccountFromConfig acc: ", acc)

	account = &dto.Account{
		Username: userInfo.Username,
	}

	resp = &dto.LoginResponseData{
		Account: account,
	}
	newAcc, _ := json.Marshal(account)
	glog.Infoln("Initialize access token")
	token, err := a.jWTAuth.InitializeToken(string(newAcc))
	if err != nil {
		glog.Errorln("InitializeToken failed:", err)
		return nil, errors.New("system error")
	}
	resp.AccessToken = token
	return resp, nil
}

func NewAuthenService(jWTAuth auth.JWTAuth, userRepo repositories.UserRepository, authConfig *config.Auth) AuthenService {
	return &authenServiceImpl{
		jWTAuth:    jWTAuth,
		userRepo:   userRepo,
		authConfig: authConfig,
	}
}

type authenServiceImpl struct {
	jWTAuth    auth.JWTAuth
	userRepo   repositories.UserRepository
	authConfig *config.Auth
}
