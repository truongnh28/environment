package v1

import (
	"github.com/gin-gonic/gin"
)

var __reportService services.reportService

//var __authenService services.AuthenService

func InitRoutes(g *gin.RouterGroup, dependencies ...interface{}) {
	for _, dependency := range dependencies {
		switch dependency.(type) {
		//case services.SongService:
		//	__songService = dependency.(services.SongService)
		//case services.AuthenService:
		//	__authenService = dependency.(services.AuthenService)
		}
	}

	reportHandler := NewReportHandler(__reportService)

	v1 := g.Group("/v1")

	// Authen
	//authenRouter := v1.Group("/authen")
	//{
	//	authenRouter.POST("login", login)
	//	authenRouter.GET("logout", logout)
	//}
	reportRouter := v1.Group("/report")
	{
		reportRouter.POST("/post_report", reportHandler.PostReport)
	}
}
