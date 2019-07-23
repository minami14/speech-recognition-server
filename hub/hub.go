package hub

import (
	"log"
	"os"

	"github.com/gorilla/websocket"
)

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	clients map[*Client]bool
	logger  *log.Logger
}

// NewHub is Hub constructed.
func NewHub() *Hub {
	return &Hub{
		clients: make(map[*Client]bool),
		logger:  log.New(os.Stdout, "WebSpeechAPIServer ", log.LstdFlags),
	}
}

// Register registers the clients.
func (h *Hub) Register(conn *websocket.Conn) {
	messageType, message, err := conn.ReadMessage()
	if err != nil {
		h.logger.Println(err)
		_ = conn.Close()
		return
	}

	if messageType != websocket.TextMessage {
		h.logger.Println("invalid message type")
		_ = conn.Close()
		return
	}

	name := string(message)
	isChrome := name == "chrome"
	client := &Client{
		hub:      h,
		conn:     conn,
		send:     make(chan []byte),
		isChrome: isChrome,
		name:     name,
	}

	h.clients[client] = true
	client.Run()
	h.logger.Println(name, "connected")
}

// SetLogger is setter for logger.
func (h *Hub) SetLogger(logger *log.Logger) {
	h.logger = logger
}

func (h *Hub) unregister(client *Client) {
	if _, ok := h.clients[client]; !ok {
		return
	}

	_ = client.conn.Close()
	delete(h.clients, client)
	h.logger.Println(client.name, "disconnected")
}

func (h *Hub) receive(message []byte) {
	h.logger.Println(string(message))
	for client := range h.clients {
		if client.isChrome {
			continue
		}

		client.send <- message
	}
}
