package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/truongnh28/environment-be/dto"
	"github.com/truongnh28/environment-be/helper"
	"github.com/truongnh28/environment-be/helper/common"
	"github.com/truongnh28/environment-be/services"
)

type UserHandlerImpl struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{
		userService: userService,
	}
}

func (s *UserHandlerImpl) GetAllUser(context *gin.Context) {
	var (
		out = &dto.UserResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	response, code := s.userService.GetAllUser()
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Users = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *UserHandlerImpl) GetUserByUserName(context *gin.Context, username string) {
	var (
		out = &dto.UserResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	response, code := s.userService.GetUserByUsername(context, username)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Users = []dto.User{response}
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}
