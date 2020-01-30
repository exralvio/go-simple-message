package services

import (
	"errors"
	"testing"
	"time"

	"github.com/exralvio/go-simple-message/cmd/http/requests"
	repoMocks "github.com/exralvio/go-simple-message/gen/mock/repositories"
	"github.com/exralvio/go-simple-message/models"
	"github.com/exralvio/go-simple-message/repositories"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("success create", func(t *testing.T) {
		mockRepo := repoMocks.NewMockMessageRepositoryInterface(mockCtrl)
		testRequest := requests.CreateMessageRequest{
			Message: "Hello World!",
		}

		model := models.CreateMessage(testRequest.Message)
		mockRepo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()

		service := &MessageService{
			repository: mockRepo,
		}
		res, err := service.Create(testRequest)

		assert.Equal(t, model.Message, res.Message)
		assert.NoError(t, err)
	})

	t.Run("failed create", func(t *testing.T) {
		mockRepo := repoMocks.NewMockMessageRepositoryInterface(mockCtrl)
		testRequest := requests.CreateMessageRequest{
			Message: "",
		}

		expected := (*models.Message)(nil)
		service := &MessageService{
			repository: mockRepo,
		}
		res, err := service.Create(testRequest)

		assert.Equal(t, expected, res)
		assert.Error(t, err)
	})

	t.Run("error database", func(t *testing.T) {
		mockRepo := repoMocks.NewMockMessageRepositoryInterface(mockCtrl)
		testRequest := requests.CreateMessageRequest{
			Message: "Hello World!",
		}

		mockRepo.EXPECT().Save(gomock.Any()).Return(errors.New("error connection")).Times(1)
		service := &MessageService{
			repository: mockRepo,
		}
		_, err := service.Create(testRequest)

		assert.Equal(t, errors.New("error connection"), err)
		assert.Error(t, err)
	})

}

func TestGetAll(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("test init service", func(t *testing.T) {
		storage := repositories.NewMessageStorage()
		NewMessageService(storage)
	})

	t.Run("success get data", func(t *testing.T) {
		repoMock := repoMocks.NewMockMessageRepositoryInterface(mockCtrl)
		var expected []models.Message

		testMessage := models.Message{
			ID:        uuid.NewV1(),
			Message:   "Hello World!",
			CreatedAt: time.Now(),
		}

		expected = append(expected, testMessage)

		repoMock.EXPECT().GetAll().Return(expected, nil).Times(1)

		service := &MessageService{
			repository: repoMock,
		}

		results, err := service.GetAll()

		assert.Equal(t, expected, results)
		assert.NoError(t, err)
	})

	t.Run("failed get data", func(t *testing.T) {
		repoMock := repoMocks.NewMockMessageRepositoryInterface(mockCtrl)
		expected := ([]models.Message)(nil)

		repoMock.EXPECT().GetAll().Return(nil, errors.New("error database")).Times(1)

		service := &MessageService{
			repository: repoMock,
		}

		results, err := service.GetAll()

		assert.Equal(t, expected, results)
		assert.Equal(t, errors.New("error database"), err)
		assert.Error(t, err)
	})
}
