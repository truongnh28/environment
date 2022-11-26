package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"os"
	"path"
	"spotify/cache"
	"spotify/client"
	"spotify/config"
	v1 "spotify/controller/v1"
	"strings"
	"time"
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
	// Init instance
	//jedis := getRedisClient()
	cldClient := client.GetCloudinaryAPI()
	//get config
	//db := getDatabaseConnector()
	// Init Repository

	// Init Service
	//memoryCache := cache.NewMemoryCache()

	resp, err := cldClient.UploadImage(context.Background(), "img.png")
	if err != nil {
		fmt.Println("err ")
		fmt.Println(err)
	}
	e, _ := json.Marshal(resp)
	fmt.Println(string(e))
	// Init w
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	//router.Use(middleware.CORSMiddleware())
	api := router.Group("/api")
	healthAPI := router.Group("/")
	healthAPI.GET("/info", getAll)
	healthAPI.GET("/health", getAll)
	healthAPI.GET("/health-check", getAll)
	healthAPI.GET("/metrics", getAll)
	healthAPI.GET("/version", getVersion)

	v1.InitRoutes(
		api,
	)
	glog.Infof("runing on port: %d ", 8080)
	err = router.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Cannot start web application with error: %v", err))
	}
}

//	func getCloudinaryClient() *cloudinary.Cloudinary {
//		const cldUrl = "cloudinary://512616158545567:mClhxuKZ9F-EsP4Kjm_s5qccdvk@dbk0cmzcb"
//		var cld, err = cloudinary.NewFromURL(cldUrl)
//		if err != nil {
//			//log.Fatalf("Failed to intialize Cloudinary, %v", err)
//			panic(fmt.Errorf("unable to connect to cloudinary: %v", err.Error()))
//		}
//		return cld
//	}
func getRedisClient() cache.RedisClient {

	if viper.GetBool("app.redis.usecluster") {
		redisClient := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    strings.Split(viper.GetString("app.redis.cluster.url"), ";"),
			Password: viper.GetString("app.redis.cluster.password"),
		})
		if err := redisClient.Ping(context.Background()).Err(); err != nil {
			panic(fmt.Errorf("unable to connect to redis cluster: %v", err.Error()))
		}
		return cache.NewRedisClusterClient(redisClient)
	} else {
		redisClient := redis.NewClient(&redis.Options{
			Addr:     viper.GetString("app.redis.single.url"),
			Password: viper.GetString("app.redis.single.password"),
		})
		if err := redisClient.Ping(context.Background()).Err(); err != nil {
			panic(fmt.Errorf("unable to connect to redis: %v", err.Error()))
		}
		return cache.NewRedisSingleClient(redisClient)
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
		// Logger: logger.Default.LogMode(logger.Info),
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

func getLDAPConfig() *config.LDAP {
	return &config.LDAP{
		Addr:        viper.GetString("app.ldap.addr"),
		UseTls:      viper.GetBool("app.ldap.useTls"),
		Username:    viper.GetString("app.ldap.username"),
		Password:    viper.GetString("app.ldap.password"),
		BaseDN:      viper.GetString("app.ldap.baseDN"),
		ObjectClass: viper.GetString("app.ldap.objectClass"),
		Timeout:     viper.GetInt64("app.ldap.timeout"),
	}
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
