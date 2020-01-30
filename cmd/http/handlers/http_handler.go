package handlers

import (
	"github.com/exralvio/go-simple-message/cmd/http/container"
	"github.com/exralvio/go-simple-message/repositories"
)

// HTTPHandler func
func HTTPHandler(App *container.AppContainer) {
	Route := App.HTTPService

	storage := repositories.NewMessageStorage()
	handler := NewMessagehandler(storage)

	Route.Static("/", "public")

	Route.POST("api/message", handler.SendMessage)
	Route.GET("api/message", handler.GetMessage)
	Route.GET("api/ws", handler.MessageWS)
}
