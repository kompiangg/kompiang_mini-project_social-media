package models

import "time"

type Notification struct {
	ID                 uint    `gorm:"primaryKey;type:BIGINT;autoIncrement"`
	TargetNotification *string `gorm:"size:36"`
	Content            string  `gorm:"size:75;type:varchar(75) NOT NULL"`
	CreatedAt          time.Time
}
