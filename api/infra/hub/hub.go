package hub

import "github.com/gorilla/websocket"

type Hub struct {
	Clients      map[*Client]bool
	Broadcast    chan []byte
	AddClient    chan *Client
	RemoveClient chan *Client
	started      bool
}

var hub = Hub{
	Broadcast:    make(chan []byte),
	AddClient:    make(chan *Client),
	RemoveClient: make(chan *Client),
	Clients:      make(map[*Client]bool),
}

func (hub *Hub) start() {
	hub.started = true
	for {
		select {
		case conn := <-hub.AddClient:
			hub.Clients[conn] = true
		case conn := <-hub.RemoveClient:
			if _, ok := hub.Clients[conn]; ok {
				delete(hub.Clients, conn)
				close(conn.Send)
			}
		case message := <-hub.Broadcast:
			for conn := range hub.Clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(hub.Clients, conn)
				}
			}
		}
	}
}

type Client struct {
	Ws   *websocket.Conn
	Send chan []byte
}

func (c *Client) Write() {
	defer func() {
		c.Ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.Ws.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (c *Client) Read() {
	defer func() {
		hub.RemoveClient <- c
		c.Ws.Close()
	}()

	for {
		_, message, err := c.Ws.ReadMessage()
		if err != nil {
			hub.RemoveClient <- c
			c.Ws.Close()
			break
		}

		hub.Broadcast <- message
	}
}

func GetHub() *Hub {
	if hub.started {
		return &hub
	}
	go hub.start()
	return &hub
}
