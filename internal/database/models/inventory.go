package models

import "time"

type Inventory struct {
	GroupID   int64  `gorm:"primaryKey"`
	UserID    int64  `gorm:"primaryKey"`
	ItemID    string `gorm:"primaryKey"`
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
