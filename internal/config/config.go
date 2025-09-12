package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DebugMode      bool
	RabbitMQURL    string
	ExchangeName   string
	QueueName      string
	Timezone       string
	TelegramToken  string
	TelegramChatID int64
}

func Load() *Config {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	debugMode, _ := strconv.ParseBool(os.Getenv("DEBUG_MODE"))
	chatID, _ := strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)

	return &Config{
		DebugMode:      debugMode,
		RabbitMQURL:    getEnv("RABBITMQ_URL", ""),
		ExchangeName:   getEnv("RABBITMQ_EXCHANGE", ""),
		QueueName:      getEnv("RABBITMQ_QUEUE", ""),
		Timezone:       getEnv("TIMEZONE", "GMT+8"),
		TelegramToken:  getEnv("TELEGRAM_BOT_TOKEN", ""),
		TelegramChatID: chatID,
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		if defaultValue == "" {
			log.Fatalf("%s environment variable is required", key)
		}
		return defaultValue
	}
	return value
}