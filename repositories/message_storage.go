package repositories

import "github.com/exralvio/go-simple-message/models"

// MessageStorage is to define message storage
type MessageStorage struct {
	Messages []models.Message
}

// NewMessageStorage is to initialize new message storage
func NewMessageStorage() *MessageStorage {
	return &MessageStorage{
		Messages: []models.Message{},
	}
}
