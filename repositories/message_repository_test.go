package repositories

import "testing"

import "github.com/exralvio/go-simple-message/models"

import "github.com/stretchr/testify/assert"

import (
	uuid "github.com/satori/go.uuid"
	"time"
) 

func TestGetAll(t *testing.T) {
	message := "Hello world!"

	storage := NewMessageStorage()
	storage.Messages = []models.Message{
		models.Message{Message: message},
	}

	t.Run("message found", func(t *testing.T) {
		repo := NewMessageRepository(storage)
		results, err := repo.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, len(storage.Messages), len(results))
	})
}

func TestSave(t *testing.T) {
	testMessage := models.Message{
		ID:        uuid.NewV1(),
		Message:   "Hello world!",
		CreatedAt: time.Now(),
	}
	storage := NewMessageStorage()

	t.Run("message saved", func(t *testing.T) {
		repo := NewMessageRepository(storage)
		err := repo.Save(testMessage)

		assert.NoError(t, err)
	})
}
