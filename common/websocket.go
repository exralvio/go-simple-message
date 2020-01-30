package common

import (
	"errors"
	"net/http"

	"github.com/exralvio/go-simple-message/models"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

// Websocket interface
type Websocket interface {
	NewConnection(w http.ResponseWriter, r *http.Request) error
	Broadcast(data *models.Message) error
}

// Ws struct
type Ws struct {
	upgrader websocket.Upgrader
	conn     *websocket.Conn
}

// SocketResponse struct
type SocketResponse struct {
	ID        uuid.UUID
	Message   string
	CreatedAt string
}

// NewWebsocket is to initialize new WS instance
func NewWebsocket() Websocket {
	wsupgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return &Ws{
		upgrader: wsupgrader,
	}
}

// NewConnection is to open connection to WS
func (s *Ws) NewConnection(w http.ResponseWriter, r *http.Request) error {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	s.conn = conn

	return nil
}

// Broadcast is to send message to opened connection
func (s *Ws) Broadcast(data *models.Message) error {
	if s.conn == nil {
		return errors.New("no listener are connected")
	}

	return s.conn.WriteJSON(SocketResponse{
		ID:        data.ID,
		Message:   data.Message,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
