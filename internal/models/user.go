package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Username      string `gorm:"primaryKey;size:50"`
	Email         string `gorm:"size:255;uniqueIndex"`
	Password      string
	UserDetail    UserDetail  `gorm:"foreignKey:Username;references:Username"`
	UserSetting   UserSetting `gorm:"foreignKey:Username;references:Username"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	RefreshTokens []RefreshToken `gorm:"foreignKey:Username;references:Username"`
	Notifications []Notification `gorm:"foreignKey:TargetNotification;references:Username"`
	Followings    []UserRelation `gorm:"foreignKey:FollowTo;references:Username"`
	Followers     []UserRelation `gorm:"foreignKey:FollowBy;references:Username"`
	Posts         []Post         `gorm:"foreignKey:PublishedBy;references:Username"`
	Comments      []Comment      `gorm:"foreignKey:CommentBy;references:Username"`
	ChatRooms     []ChatRoomUser `gorm:"foreignKey:SentBy;references:Username"`
	Messages      []Message      `gorm:"foreignKey:SentBy;references:Username"`
}
