package config

import (
	"log"
	"os"
)

type Config struct {
	Port            string
	TelegramToken   string
	ChatID          string
	ControlBotToken string
	RedisURL        string
	APIKey          string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	telegramToken := os.Getenv("TELEGRAM_BOT_NOTIF_TOKEN")
	if telegramToken == "" {
		log.Println("WARNING: TELEGRAM_BOT_NOTIF_TOKEN is not set")
	}

	chatID := os.Getenv("TELEGRAM_CHAT_ID")
	if chatID == "" {
		log.Println("WARNING: TELEGRAM_CHAT_ID is not set")
	}

	controlBotToken := os.Getenv("TELEGRAM_BOT_CONTROL_TOKEN")
	if controlBotToken == "" {
		log.Println("WARNING: TELEGRAM_BOT_CONTROL_TOKEN is not set")
	}

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Println("WARNING: REDIS_URL is not set. Defaulting to 'redis://localhost:6379'")
		redisURL = "redis://localhost:6379"
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Println("WARNING: API_KEY is not set. Defaulting to 'secret'")
		apiKey = "secret"
	}

	return &Config{
		Port:            port,
		TelegramToken:   telegramToken,
		ChatID:          chatID,
		ControlBotToken: controlBotToken,
		RedisURL:        redisURL,
		APIKey:          apiKey,
	}
}
