package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID          string  `gorm:"primaryKey;size:36;autoIncrement"`
	PublishedBy string  `gorm:"size:50;type:varchar(50) NOT NULL"`
	RepostID    *string `gorm:"size:36"`
	Content     string  `gorm:"size:255"`
	PictureURL  *string
	VideoURL    *string
	Posts       []Post    `gorm:"foreignKey:RepostID;references:ID"`
	Comments    []Comment `gorm:"foreignKey:PostID;references:ID"`
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
