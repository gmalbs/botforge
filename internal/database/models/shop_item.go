package models

import "time"

type ShopItem struct {
	ID         string `gorm:"primaryKey"`
	Key        string
	GroupID    int64 `gorm:"index"`
	Name       string
	Descrption string
	PriceCoins int
	Stock      *int
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
