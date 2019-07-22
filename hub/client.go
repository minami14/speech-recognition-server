package hub

import (
	"github.com/gorilla/websocket"
)

// Client is a middleman between the connection and the hub.
type Client struct {
	hub      *Hub
	conn     *websocket.Conn
	send     chan []byte
	isChrome bool
	name     string
}

// Run is provides backend synchronize goroutine.
func (c *Client) Run() {
	if c.isChrome {
		go func() {
			close(c.send)
			for {
				messageType, message, err := c.conn.ReadMessage()
				if err != nil {
					c.hub.logger.Println(err)
					c.hub.unregister(c)
					return
				}

				if messageType != websocket.TextMessage {
					c.hub.logger.Println("invalid message type")
					c.hub.unregister(c)
					return
				}

				c.hub.receive(message)
			}
		}()
	} else {
		go func() {
			for {
				message := <-c.send
				if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
					c.hub.logger.Println(err)
					c.hub.unregister(c)
				}
			}
		}()
	}
}
