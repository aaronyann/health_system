package models

import "time"

type Notification struct {
	NotificationID int       `gorm:"primaryKey" json:"notification_id"`
	UserID         uint      `json:"user_id"`
	Message        string    `json:"message"`
	ReadStatus     bool      `json:"read_status"`
	CreatedAt      time.Time `json:"created_at"`
}
