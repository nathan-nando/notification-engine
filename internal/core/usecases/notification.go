package usecases

import (
	"fmt"
	"notification-engine/internal/core/domain"
	"notification-engine/internal/core/ports"
)

// NotificationService is the driving port/usecase interface for notifications.
type NotificationService interface {
	ProcessNotification(level, source, message string) error
}

type notificationService struct {
	messenger ports.Messenger
}

// NewNotificationService creates a new notification service.
func NewNotificationService(messenger ports.Messenger) NotificationService {
	return &notificationService{
		messenger: messenger,
	}
}

// ProcessNotification implements the business logic for processing and routing a notification.
func (s *notificationService) ProcessNotification(level, source, message string) error {
	// Create the domain entity
	notification := domain.NewNotification(level, source, message)

	// Here we could add business logic, e.g., filtering based on level, formatting, etc.
	// For now, we just pass it to the messenger port.
	err := s.messenger.SendMessage(notification)
	if err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}

	return nil
}
