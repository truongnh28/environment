package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/truongnh28/environment-be/services"
)

//var __reportService services.reportService

//var __authenService services.AuthenService

var __userService services.UserService

func InitRoutes(g *gin.RouterGroup, dependencies ...interface{}) {
	for _, dependency := range dependencies {
		switch dependency.(type) {
		case services.UserService:
			__userService = dependency.(services.UserService)
			//case services.AuthenService:
			//	__authenService = dependency.(services.AuthenService)
		}
	}

	//reportHandler := NewReportHandler(__userService])

	userHandler := NewUserHandler(__userService)

	v1 := g.Group("/v1")

	userRouter := v1.Group("/user")
	{
		userRouter.GET("get_all", userHandler.GetAllUser)
	}
}
