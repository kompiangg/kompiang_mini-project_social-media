package models

import "time"

type Message struct {
	ID        uint64 `gorm:"primaryKey;type:BIGINT AUTO_INCREMENT"`
	GUID      string `gorm:"size:36;type:varchar(36) NOT NULL"`
	SentBy    string `gorm:"size:50;type:varchar(50) NOT NULL"`
	Content   string `gorm:"type:text NOT NULL"`
	CreatedAt time.Time
}
