package models

import "time"

type Comment struct {
	ID            string  `gorm:"primaryKey;size:36;autoIncrement"`
	CommentBy     string  `gorm:"size:50;type:varchar(50) NOT NULL"`
	PostID        *string `gorm:"size:36"`
	Content       string  `gorm:"size:255"`
	CommentID     *string
	PictureURL    string
	VideoURL      string
	ChildComments []Comment `gorm:"foreignKey:CommentID;references:ID"`
	CreatedAt     time.Time
}
