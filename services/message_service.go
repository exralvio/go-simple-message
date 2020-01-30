package services

import (
	"errors"
	"github.com/exralvio/go-simple-message/cmd/http/requests"
	"github.com/exralvio/go-simple-message/models"
	"github.com/exralvio/go-simple-message/repositories"
)

// MessageServiceInterface interface
type MessageServiceInterface interface {
	Create(message requests.CreateMessageRequest) (*models.Message, error)
	GetAll() ([]models.Message, error)
}

// MessageService func
type MessageService struct {
	repository repositories.MessageRepositoryInterface
}

// NewMessageService func
func NewMessageService(storage *repositories.MessageStorage) MessageServiceInterface {
	service := &MessageService{
		repository: repositories.NewMessageRepository(storage),
	}

	return service
}

// Create func
func (service *MessageService) Create(message requests.CreateMessageRequest) (*models.Message, error) {
	if message.Message == "" {
		return nil, errors.New("message is empty")
	}

	model := models.CreateMessage(message.Message)
	if err := service.repository.Save(*model); err != nil {
		return nil, err
	}

	return model, nil
}

// GetAll func
func (service *MessageService) GetAll() ([]models.Message, error) {
	models, err := service.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return models, nil
}
