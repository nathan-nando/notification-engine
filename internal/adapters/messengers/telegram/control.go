package telegram

import (
	"context"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"notification-engine/internal/adapters/broker"
)

type ControlBot struct {
	bot    *tgbotapi.BotAPI
	broker *broker.RedisBroker
}

func NewControlBot(token string, redisBroker *broker.RedisBroker) (*ControlBot, error) {
	if token == "" {
		log.Println("Control bot token is empty, skipping initialization")
		return nil, nil
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	log.Printf("Control Bot authorized on account %s", bot.Self.UserName)

	return &ControlBot{
		bot:    bot,
		broker: redisBroker,
	}, nil
}

func (c *ControlBot) Start(ctx context.Context) {
	// Subscribe to responses
	c.broker.SubscribeControlResponses(ctx, func(resp broker.ControlResponse) {
		msg := tgbotapi.NewMessage(resp.ChatID, resp.Message)
		msg.ParseMode = "Markdown"
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("Failed to send control response to chat %d: %v", resp.ChatID, err)
		}
	})

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.bot.GetUpdatesChan(u)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case update := <-updates:
				if update.Message == nil { // ignore any non-Message updates
					continue
				}

				if !update.Message.IsCommand() {
					continue
				}

				cmd := "/" + update.Message.Command()
				argsStr := update.Message.CommandArguments()
				var args []string
				if argsStr != "" {
					args = strings.Fields(argsStr)
				}

				req := broker.ControlRequest{
					ChatID:  update.Message.Chat.ID,
					Command: cmd,
					Args:    args,
				}

				if err := c.broker.PublishControlRequest(ctx, req); err != nil {
					log.Printf("Failed to publish control request: %v", err)
				}
			}
		}
	}()
}
