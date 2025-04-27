package services

import (
	"backend/models"
	"fmt"
)

// GetNotifications 获取指定用户的通知列表
func GetNotifications(userID uint) ([]models.Notification, error) {
	var notifications []models.Notification
	if err := DB.Where("user_id = ?", userID).Find(&notifications).Error; err != nil {
		return nil, fmt.Errorf("查询通知失败: %v", err)
	}
	return notifications, nil
}

// MarkNotificationRead 标记通知为已读
func MarkNotificationRead(notificationID uint) error {
	var notification models.Notification
	if err := DB.First(&notification, notificationID).Error; err != nil {
		return fmt.Errorf("查询通知失败: %v", err)
	}
	notification.ReadStatus = true
	if err := DB.Save(&notification).Error; err != nil {
		return fmt.Errorf("更新通知状态失败: %v", err)
	}
	return nil
}
