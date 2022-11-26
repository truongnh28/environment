package v1

import (
	"encoding/json"
	"fmt"
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

func (s *UserHandlerImpl) GetUserByUserName(context *gin.Context) {
	var (
		out = &dto.UserResponse{}
	)
	username, ok := context.Params.Get("username")
	if !ok {
		helper.BuildResponseByReturnCode(out, common.Fail, common.InvalidRequest)
		return
	}
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

func (s *UserHandlerImpl) Login(context *gin.Context) {
	var (
		out = &dto.UserResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	var body dto.LoginRequest
	if err := json.NewDecoder(context.Request.Body).Decode(&body); err != nil {
		helper.BuildResponseByReturnCode(out, common.Fail, common.UsernameOrPasswordIncorrect)
		return
	}
	response, code := s.userService.Login(context, body.Username, body.Password)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, common.UsernameOrPasswordIncorrect)
		return
	}
	out.Users = []dto.User{response}
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *UserHandlerImpl) Register(context *gin.Context) {
	var (
		out = &dto.UserResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	var body dto.User
	if err := json.NewDecoder(context.Request.Body).Decode(&body); err != nil {
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	response, code := s.userService.Register(context, body)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	out.Users = []dto.User{response}
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *UserHandlerImpl) ChangeInfo(context *gin.Context) {
	var (
		out = &dto.UserResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	m := make(map[string]interface{})
	if err := json.NewDecoder(context.Request.Body).Decode(&m); err != nil {
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	fmt.Println("mapData: ", m)
	response, code := s.userService.ChangeInfo(context, m)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	out.Users = []dto.User{response}
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}
