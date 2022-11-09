package models

import "time"

type UserSetting struct {
	ID              uint   `gorm:"primaryKey"`
	Username        string `gorm:"size:50"`
	IsVerifiedEmail bool
	IsDeactivate    bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
