package api

import (
	"net/http"

	"github.com/gmalbs/botforge/internal/database"
	"github.com/gmalbs/botforge/internal/database/models"

	"github.com/gin-gonic/gin"
)

func StartServer(port string) {
	r := gin.Default()

	// Basic API for Mini App
	r.GET("/api/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user models.User
		if err := database.DB.Where("user_id = ?", id).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	r.GET("/api/status", func(c *gin.Context) {
		var settings models.BotSettings
		database.DB.First(&settings)
		c.JSON(http.StatusOK, gin.H{
			"maintenance": settings.MaintenanceMode,
			"status":      "online",
		})
	})

	// Serve static files for Mini App if needed
	// r.Static("/miniapp", "./web/dist")

	go r.Run(":" + port)
}
