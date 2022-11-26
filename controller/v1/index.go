package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/truongnh28/environment-be/services"
)

var __userService services.UserService
var __reportService services.ReportService
var __authenService services.AuthenService

func InitRoutes(g *gin.RouterGroup, dependencies ...interface{}) {
	for _, dependency := range dependencies {
		switch dependency.(type) {
		case services.UserService:
			__userService = dependency.(services.UserService)
		case services.ReportService:
			__reportService = dependency.(services.ReportService)
		case services.AuthenService:
			__authenService = dependency.(services.AuthenService)
		}
	}

	userHandler := NewUserHandler(__userService)
	reportHandler := NewReportHandler(__reportService)
	v1 := g.Group("/v1")
	authenRouter := v1.Group("/authen")
	{
		authenRouter.POST("login", login)
		authenRouter.GET("logout", logout)
	}
	userRouter := v1.Group("/user")
	{
		userRouter.GET("get_all", userHandler.GetAllUser)
		userRouter.POST("login", userHandler.Login)
		userRouter.POST("register", userHandler.Register)
		userRouter.POST("change_info", userHandler.ChangeInfo)
	}
	reportRouter := v1.Group("/report")
	{
		reportRouter.GET("map_list", reportHandler.MapResponseList)
		reportRouter.POST("create", reportHandler.CreateReport)
		reportRouter.GET("get_by_id", reportHandler.GetReportByID)
		reportRouter.POST("list", reportHandler.ListReport)
		reportRouter.POST("upload_file", reportHandler.UploadImage)
		reportRouter.GET("get_all", reportHandler.GetAllReport)
		reportRouter.POST("update", reportHandler.UpdateReport)
		reportRouter.GET("top", reportHandler.TopResolver)
	}
}
