package ports

import "notification-engine/internal/core/domain"

// Messenger is the driven port for sending notifications to external platforms.
type Messenger interface {
	SendMessage(notification *domain.Notification) error
}
