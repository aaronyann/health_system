package controllers

import (
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNotifications 获取用户通知列表
func GetNotifications(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.ResponseError(c, http.StatusUnauthorized, "未登录")
		return
	}
	notifications, err := services.GetNotifications(userID.(uint))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "获取通知失败")
		return
	}
	utils.ResponseSuccess(c, notifications)
}

// MarkNotificationRead 标记通知为已读
func MarkNotificationRead(c *gin.Context) {
	var req struct {
		NotificationID uint `json:"notification_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "参数错误")
		return
	}
	err := services.MarkNotificationRead(req.NotificationID)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "标记失败: "+err.Error())
		return
	}
	utils.ResponseSuccess(c, gin.H{"message": "通知已标记为已读"})
}
