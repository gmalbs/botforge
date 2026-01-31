package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	TelegramBotToken string
	Database         string
	AppID            int64
	AppHash          string
	SecreteKey       string
	WebAppURL        string
	OwnerID          int64
)

func init() {
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("⚠️  .env não encontrado — usando variáveis de ambiente do container")
		}
	}
	fmt.Println(os.Getenv("DATABASE_FILE"))

	TelegramBotToken = mustGetEnv("BOT_TOKEN")
	AppID = mustGetEnvInt64("APP_ID")
	AppHash = os.Getenv("APP_HASH")
	Database = os.Getenv("DATABASE")
	OwnerID = mustGetEnvInt64("OWNER_ID")
	SecreteKey = mustGetEnv("SECRET_KEY")
	WebAppURL = mustGetEnv("WEBAPP_URL")
}

func mustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("Environment variable %s is required", key)
	}
	return v
}

func mustGetEnvInt64(key string) int64 {
	v := mustGetEnv(key)
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Fatalf("Environment variable %s must be an integer: %v", key, err)
	}
	return n
}
