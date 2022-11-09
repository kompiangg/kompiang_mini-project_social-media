package models

import "time"

type UserRelation struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;type:BIGINT"`
	FollowBy  string `gorm:"size:50;type:varchar(50) NOT NULL;index:idx_user_relations"`
	FollowTo  string `gorm:"size:50;type:varchar(50) NOT NULL;index:idx_user_relations"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
