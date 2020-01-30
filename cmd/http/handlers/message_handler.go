package handlers

import (
	"log"
	"net/http"

	"github.com/exralvio/go-simple-message/cmd/http/requests"
	"github.com/exralvio/go-simple-message/common"
	"github.com/exralvio/go-simple-message/repositories"
	"github.com/exralvio/go-simple-message/services"
	"github.com/labstack/echo"
)

// MessageHandler struct
type MessageHandler struct {
	service   services.MessageServiceInterface
	websocket common.Websocket
}

// NewMessagehandler func
func NewMessagehandler(storage *repositories.MessageStorage, websocket common.Websocket) *MessageHandler {
	handler := &MessageHandler{
		service:   services.NewMessageService(storage),
		websocket: websocket,
	}

	return handler
}

// SendMessage func
func (handler *MessageHandler) SendMessage(c echo.Context) error {
	var request requests.CreateMessageRequest

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid request",
		})

		return nil
	}

	result, err := handler.service.Create(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})

		return nil
	}

	if err := handler.websocket.Broadcast(result); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})

		return nil
	}

	return c.JSON(http.StatusCreated, result)
}

// GetMessage func
func (handler *MessageHandler) GetMessage(c echo.Context) error {
	results, err := handler.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})

		return nil
	}

	response := map[string]interface{}{
		"data": results,
	}

	return c.JSON(http.StatusOK, response)
}

// MessageWS func
func (handler *MessageHandler) MessageWS(c echo.Context) error {
	err := handler.websocket.NewConnection(c.Response(), c.Request())

	if err != nil {
		log.Printf("websocket terminated with error: %s", err)
	}

	return nil
}
