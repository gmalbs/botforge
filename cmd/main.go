package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gmalbs/botforge/api"
	"github.com/gmalbs/botforge/internal/bot"
	"github.com/gmalbs/botforge/internal/config"
	"github.com/gmalbs/botforge/internal/database"

	// Import modules to trigger init()
	"github.com/gmalbs/botforge/internal/modules"
)

func main() {
	// 1. Load Configs
	err := config.LoadMessages("assets/messages.yml")
	if err != nil {
		log.Fatal("Failed to load messages:", err)
	}

	// 2. Init Database
	database.InitDB()

	// Trigger module registration
	modules.Load()

	// 3. Start API Server
	api.StartServer("8080")

	// 4. Init Bot

	client, err := bot.InitBot(config.TelegramBotToken)
	if err != nil {
		log.Fatal("Failed to initialize bot:", err)
	}

	log.Println("Bot is running...")

	// Keep running
	client.Idle()

	// Handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down...")
}
