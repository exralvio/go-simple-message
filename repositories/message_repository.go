package repositories

import "github.com/exralvio/go-simple-message/models"

// MessageRepositoryInterface is to initialize interface
type MessageRepositoryInterface interface {
	GetAll() ([]models.Message, error)
	Save(data models.Message) error
}

// MessageRepository struct
type MessageRepository struct {
	storage *MessageStorage
}

// NewMessageRepository func
func NewMessageRepository(storage *MessageStorage) MessageRepositoryInterface {
	repository := &MessageRepository{
		storage: storage,
	}

	return repository
}

// GetAll func
func (repo *MessageRepository) GetAll() ([]models.Message, error) {
	var results []models.Message

	for _, item := range repo.storage.Messages {
		results = append(results, item)
	}

	return results, nil
}

// Save func
func (repo *MessageRepository) Save(data models.Message) error {
	repo.storage.Messages = append(repo.storage.Messages, data)

	return nil
}
