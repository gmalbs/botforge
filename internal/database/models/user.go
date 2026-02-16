package models

import (
	"gorm.io/gorm"
)

type User struct {
	UserID    int64 `gorm:"uniqueIndex"`
	FirstName string
	Username  string
	IsAdmin   bool `gorm:"default:false"`
}

type BotSettings struct {
	gorm.Model
	MaintenanceMode bool `gorm:"default:false"`
}
