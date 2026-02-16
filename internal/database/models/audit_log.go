package models

import "time"

type AuditLog struct {
	ID        string `gorm:"primaryKey"`
	BotID     int64  `gorm:"index"`
	GroupID   *int64 `gorm:"index"`
	UserID    *int64 `gorm:"index"`
	Action    string
	Metadata  string `gorm:"type:jsonb"`
	CreatedAt time.Time
}
