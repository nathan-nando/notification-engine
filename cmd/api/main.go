package main

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"notification-engine/config"
	_ "notification-engine/docs" // Swagger docs
	"notification-engine/internal/adapters/broker"
	"notification-engine/internal/adapters/handlers/http"
	"notification-engine/internal/adapters/messengers/telegram"
	"notification-engine/internal/core/usecases"
)

// @title Notification Engine API
// @version 1.0
// @description This is the API for the Quant Notification Engine.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
func main() {
	cfg := config.LoadConfig()

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize Redis Broker
	redisBroker, err := broker.NewRedisBroker(cfg.RedisURL)
	if err != nil {
		log.Fatalf("Failed to initialize redis broker: %v", err)
	}
	log.Println("Successfully connected to Redis PubSub")

	// Initialize Ports and Adapters
	var messenger usecases.NotificationService
	
	if cfg.TelegramToken != "" && cfg.ChatID != "" {
		tgAdapter, err := telegram.NewTelegramAdapter(cfg.TelegramToken, cfg.ChatID)
		if err != nil {
			log.Fatalf("Failed to initialize telegram adapter: %v", err)
		}
		messenger = usecases.NewNotificationService(tgAdapter)
	} else {
		log.Println("WARNING: Running without Telegram bot configured. Notifications will fail or be logged.")
	}

	// Initialize Control Bot
	if cfg.ControlBotToken != "" {
		controlBot, err := telegram.NewControlBot(cfg.ControlBotToken, redisBroker)
		if err != nil {
			log.Fatalf("Failed to initialize control bot: %v", err)
		}
		if controlBot != nil {
			ctx := context.Background()
			controlBot.Start(ctx)
			log.Println("Control Bot listener started.")
		}
	} else {
		log.Println("WARNING: CONTROL_BOT_TOKEN not configured. Two-way commands disabled.")
	}

	handler := http.NewNotificationHandler(messenger)

	// API Key Middleware for protected routes
	apiKeyMiddleware := middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:X-API-Key",
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == cfg.APIKey, nil
		},
	})

	// Routes
	v1 := e.Group("/api/v1")
	v1.Use(apiKeyMiddleware)
	v1.POST("/notify", handler.SendNotification)

	// Swagger Docs (Public)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	log.Printf("Starting Notification Engine on port %s", cfg.Port)
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
