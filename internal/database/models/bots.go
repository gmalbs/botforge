package models

import "time"

type Bot struct {
	ID        int64 `gorm:"primaryKey"`
	Name      string
	Username  string
	TokenHash string `gorm:"uniqueIndex"`
	OwnerID   int64
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
