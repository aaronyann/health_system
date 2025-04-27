package router

import (
	"backend/controllers"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// 认证路由组
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", controllers.Login)
		authGroup.POST("/register", controllers.Register)
	}

	// API 路由组（需要认证）
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// 用户信息
		api.GET("/profile", controllers.GetProfile)
		api.PUT("/profile", controllers.UpdateProfile)
		// 健康档案
		api.GET("/health/records", controllers.GetHealthRecords)
		api.POST("/health/records", controllers.CreateHealthRecord)
		api.GET("/health/records/:id", controllers.GetHealthRecordByID)
		// 预约
		api.GET("/appointments", controllers.GetAppointments)
		api.POST("/appointments", controllers.CreateAppointment)
		api.GET("/appointments/:id", controllers.GetAppointmentByID)
		// 处方
		api.GET("/prescriptions", controllers.GetPrescriptionList)
		api.GET("/prescriptions/:id", controllers.GetPrescriptionDetail)
		api.POST("/prescriptions", controllers.CreatePrescription)
		// 检验报告
		api.GET("/lab/reports", controllers.GetLabReports)
		api.GET("/lab/reports/:id", controllers.GetLabReportByID)
		api.POST("/lab/reports", controllers.CreateLabReport)
		// 保险信息
		api.GET("/insurance", controllers.GetInsuranceInfo)
		api.PUT("/insurance", controllers.UpdateInsuranceInfo)
		// 通知
		api.GET("/notifications", controllers.GetNotifications)
		api.POST("/notifications/read", controllers.MarkNotificationRead)
	}
}
