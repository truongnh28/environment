package main

import (
	"flag"
	"fmt"
	"github.com/truongnh28/environment-be/config"
	"github.com/truongnh28/environment-be/helper"
	"github.com/truongnh28/environment-be/middleware"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	v1 "github.com/truongnh28/environment-be/controller/v1"
	"github.com/truongnh28/environment-be/repositories"
	"github.com/truongnh28/environment-be/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()
	configPath, configFile := extractConfigPath()
	viper.AddConfigPath(configPath)
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")
	viper.SetConfigName(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	db := getDatabaseConnector()

	userRepo := repositories.NewUserRepository(db)
	reportRepo := repositories.NewReportRepository(db)

	userServices := services.NewUserService(userRepo)
	reportService := services.NewReportService(reportRepo)
	authenService := services.NewAuthenService(helper.GetJWTInstance(), userRepo, config.AuthConfig())

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	api := router.Group("/api")
	healthAPI := router.Group("/")
	healthAPI.GET("/info", getAll)
	healthAPI.GET("/health", getAll)
	healthAPI.GET("/health-check", getAll)
	healthAPI.GET("/metrics", getAll)
	healthAPI.GET("/version", getVersion)

	v1.InitRoutes(
		api,
		userServices,
		reportService,
		authenService,
	)
	glog.Infof("runing on port: %d ", 8080)
	err = router.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Cannot start web application with error: %v", err))
	}
}

func getDatabaseConnector() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("app.database.username"),
		viper.GetString("app.database.password"),
		viper.GetString("app.database.address"),
		viper.GetInt("app.database.port"),
		viper.GetString("app.database.database-name"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(fmt.Errorf("failed to connect database, error: %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(viper.GetInt("app.database.max-open-conns"))
	sqlDB.SetMaxIdleConns(viper.GetInt("app.database.max-idle-conns"))
	sqlDB.SetConnMaxLifetime(time.Duration(viper.GetInt("app.database.conn-max-lifetime")) * time.Minute)

	return db
}

func getAll(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func getVersion(c *gin.Context) {
	c.JSON(http.StatusOK, os.Getenv("image_tag"))
}

func extractConfigPath() (string, string) {
	var (
		defaultConfig = "config/local.yml"
		cp            = os.Getenv("CONFIG_PATH")
	)
	if len(cp) > 0 {
		defaultConfig = cp
	}

	configPath, configFile := path.Split(defaultConfig)
	ext := path.Ext(configFile)
	if len(ext) > 0 {
		configFile = configFile[:len(configFile)-len(ext)]
	}
	return configPath, configFile
}
