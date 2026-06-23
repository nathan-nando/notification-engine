package config

import (
	"log"
	"os"
)

type Config struct {
	Port          string
	TelegramToken string
	ChatID        string
	APIKey        string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	telegramToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if telegramToken == "" {
		log.Println("WARNING: TELEGRAM_BOT_TOKEN is not set")
	}

	chatID := os.Getenv("TELEGRAM_CHAT_ID")
	if chatID == "" {
		log.Println("WARNING: TELEGRAM_CHAT_ID is not set")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Println("WARNING: API_KEY is not set. Defaulting to 'secret'")
		apiKey = "secret"
	}

	return &Config{
		Port:          port,
		TelegramToken: telegramToken,
		ChatID:        chatID,
		APIKey:        apiKey,
	}
}
