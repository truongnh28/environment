package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/truongnh28/environment-be/services"
)

var __userService services.UserService
var __reportService services.ReportService

func InitRoutes(g *gin.RouterGroup, dependencies ...interface{}) {
	for _, dependency := range dependencies {
		switch dependency.(type) {
		case services.UserService:
			__userService = dependency.(services.UserService)
		case services.ReportService:
			__reportService = dependency.(services.ReportService)
		}
	}

	//reportHandler := NewReportHandler(__userService])

	userHandler := NewUserHandler(__userService)
	reportHandler := NewReportHandler(__reportService)
	v1 := g.Group("/v1")

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
	}
}
