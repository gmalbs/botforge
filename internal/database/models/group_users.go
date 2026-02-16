package models

import "time"

type GroupUser struct {
	GroupID      int64 `gorm:"primaryKey"`
	UserID       int64 `gorm:"primaryKey"`
	MessageCount int
	Coins        int64
	LastRewardAt *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
