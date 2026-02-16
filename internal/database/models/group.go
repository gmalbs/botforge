package models

import "time"

type Group struct {
	ID        int64 `gorm:"primaryKey"`
	BotID     int64 `gorm:"index"`
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
