package models

import (
	"time"
)

type RefreshToken struct {
	ID             uint `gorm:"primarykey;autoIncrement"`
	Username       string
	Token          string
	CreatedAt      time.Time
	ExpirationDate time.Time
}
