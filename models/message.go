package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Message model
type Message struct {
	ID        uuid.UUID
	Message   string
	CreatedAt time.Time
}

// CreateMessage is to initialize new message
func CreateMessage(message string) *Message {
	return &Message{
		ID:        uuid.NewV1(),
		Message:   message,
		CreatedAt: time.Now().Local(),
	}
}
