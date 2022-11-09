package dto

import "time"

type AdminNotificationRequest struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type NotificationResponse struct {
	Content            string    `json:"content"`
	TargetNotification *string   `gorm:"target_notification"`
	CreatedAt          time.Time `json:"created_at"`
}

type NotificationsResponse []NotificationResponse
