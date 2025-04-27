package main

import (
	"backend/config"
	"backend/models"
	"backend/router"
	"backend/services"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := config.AppConfig.Database.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败: ", err)
	}
	// 自动迁移数据库模型
	db.AutoMigrate(&models.User{}, &models.HealthRecord{}, &models.Appointment{}, &models.Prescription{}, &models.LabReport{}, &models.InsuranceInfo{}, &models.Notification{})
	return db
}

func main() {
	// 加载配置
	config.LoadConfig("config/config.yaml")

	// 初始化数据库
	db := initDB()
	services.DB = db // 将数据库赋值给服务层全局变量

	// 创建 Gin 实例
	r := gin.Default()
	// 添加全局中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 注册路由
	router.RegisterRoutes(r)

	// 启动 HTTP 服务器
	srv := &http.Server{
		Addr:         config.AppConfig.Server.Port,
		Handler:      r,
		ReadTimeout:  config.AppConfig.Server.ReadTimeout,
		WriteTimeout: config.AppConfig.Server.WriteTimeout,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("启动服务器失败: ", err)
	}
}
