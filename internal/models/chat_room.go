package models

import "time"

type ChatRoom struct {
	GUID         string `gorm:"primaryKey;size:36"`
	CreatedAt    time.Time
	ChatRoomUser []ChatRoomUser `gorm:"foreignKey:GUID;references:GUID"`
	Messages     []Message      `gorm:"foreignKey:GUID;references:GUID"`
}

type ChatRoomUser struct {
	ID     uint64 `gorm:"primaryKey;type:BIGINT AUTO_INCREMENT"`
	GUID   string `gorm:"size:36;type:varchar(36) NOT NULL"`
	SentBy string `gorm:"size:50;type:varchar(36) NOT NULL"`
}
