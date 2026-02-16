package models

import "time"

type RewardRule struct {
	ID               string `gorm:"primaryKey"`
	GroupID          int64  `gorm:"index"`
	MessagesRequired int
	CoinsReward      int
	CooldownType     int
	CooldownSeconds  *int
	IsActive         bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
