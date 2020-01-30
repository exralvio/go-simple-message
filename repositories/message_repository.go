package repositories

import "github.com/exralvio/go-simple-message/models"

// MessageRepositoryInterface is to initialize repository interface
type MessageRepositoryInterface interface {
	GetAll() ([]models.Message, error)
	Save(data models.Message) error
}

// MessageRepository struct
type MessageRepository struct {
	storage *MessageStorage
}

// NewMessageRepository is to initialize repository
func NewMessageRepository(storage *MessageStorage) MessageRepositoryInterface {
	repository := &MessageRepository{
		storage: storage,
	}

	return repository
}

// GetAll is to retrieve all message
func (repo *MessageRepository) GetAll() ([]models.Message, error) {
	var results []models.Message

	for _, item := range repo.storage.Messages {
		results = append(results, item)
	}

	return results, nil
}

// Save is to save new message
func (repo *MessageRepository) Save(data models.Message) error {
	repo.storage.Messages = append(repo.storage.Messages, data)

	return nil
}
