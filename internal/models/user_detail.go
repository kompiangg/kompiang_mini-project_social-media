package models

import "time"

type UserDetail struct {
	ID                uint   `gorm:"primaryKey"`
	Username          string `gorm:"size:50"`
	DisplayName       string `gorm:"size:100;type:varchar(100) NOT NULL"`
	ProfilePictureURL string
	Bio               string     `gorm:"size:255"`
	BirthDate         *time.Time `gorm:"type:DATE"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
