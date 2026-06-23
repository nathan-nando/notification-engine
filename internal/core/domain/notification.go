package domain

import "time"

// Notification represents the core business entity for a message to be sent.
type Notification struct {
	Level     string    `json:"level"` // e.g., "INFO", "WARNING", "ERROR"
	Source    string    `json:"source"` // e.g., "Quant Engine", "Train Engine"
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// NewNotification creates a new Notification entity with current timestamp if not provided.
func NewNotification(level, source, message string) *Notification {
	return &Notification{
		Level:     level,
		Source:    source,
		Message:   message,
		Timestamp: time.Now(),
	}
}
