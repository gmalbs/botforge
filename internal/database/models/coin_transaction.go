package models

import "time"

type CoinTransaction struct {
	ID        string `gorm:"primaryKey"`
	GroupID   int64  `gorm:"index"`
	UserID    int64  `gorm:"index"`
	Amount    int
	Reason    string
	CreatedAt time.Time
}
