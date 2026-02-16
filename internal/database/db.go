package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/gmalbs/botforge/internal/config"
	"github.com/gmalbs/botforge/internal/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(postgres.Open(config.Database), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate
	err = db.AutoMigrate(
		&models.User{},
		&models.BotSettings{},

		&models.Bot{},
		&models.Group{},
		&models.GroupUser{},
		&models.RewardRule{},
		&models.ShopItem{},
		&models.Inventory{},
		&models.CoinTransaction{},
		&models.AuditLog{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize settings if not exists
	var settings models.BotSettings
	if err := db.First(&settings).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&models.BotSettings{
				MaintenanceMode: false,
			})
		} else {
			log.Fatal("Failed to read bot settings:", err)
		}
	}

	DB = db
	fmt.Println("Database connected and migrated.")
}
