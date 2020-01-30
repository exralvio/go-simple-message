package common

import (
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

// WebSocketConnection struct
type WebSocketConnection struct {
	*websocket.Conn
	ID string
}

var currentConnection WebSocketConnection
var activeConnections = make([]*WebSocketConnection, 0)

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

	connID := r.URL.Query().Get("id")
	currentConn := WebSocketConnection{Conn: conn, ID: connID}
	currentConnection = currentConn
	activeConnections = append(activeConnections, &currentConn)

	return nil
}

// Broadcast is to send message to opened connection
func (s *Ws) Broadcast(data *models.Message) error {
	for _, conn := range activeConnections {
		err := conn.WriteJSON(SocketResponse{
			ID:        data.ID,
			Message:   data.Message,
			CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
		})

		if err != nil {
			return err
		}
	}

	return nil
}
