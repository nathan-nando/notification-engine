package telegram

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"notification-engine/internal/core/domain"
	"notification-engine/internal/core/ports"
)

type telegramAdapter struct {
	bot    *tgbotapi.BotAPI
	chatID int64
}

// NewTelegramAdapter creates a new Telegram messenger adapter.
func NewTelegramAdapter(token, chatIDStr string) (ports.Messenger, error) {
	if token == "" {
		return nil, fmt.Errorf("telegram token is empty")
	}

	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid chat ID: %w", err)
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create telegram bot: %w", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &telegramAdapter{
		bot:    bot,
		chatID: chatID,
	}, nil
}

// SendMessage implements the Messenger interface.
func (t *telegramAdapter) SendMessage(notification *domain.Notification) error {
	// Format the message
	text := notification.Message

	msg := tgbotapi.NewMessage(t.chatID, text)
	msg.ParseMode = "Markdown"

	_, err := t.bot.Send(msg)
	return err
}
