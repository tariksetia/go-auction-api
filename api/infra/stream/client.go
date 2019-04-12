package stream

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"time"

)

type Client struct {
	Ws   *websocket.Conn
	Send chan []byte
}

func (c *Client) Write(hub *Hub) {
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

func (c *Client) Read(hub *Hub) {
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

		var socketMessage = SocketMessage{}
		err = json.Unmarshal([]byte(message), &socketMessage)
		if err != nil {
			hub.RemoveClient <- c
			c.Ws.Close()
			break
		}

		hubMessage := HubMessage{
			Data: &socketMessage,
			Client: c,
			Hub: hub,
		}

		hub.Incoming <- &hubMessage

	}
}

func (c *Client) CloseWithError(err string) {
	//msg := SocketError{
	//	Error: err,
	//}
	c.Ws.WriteMessage(websocket.TextMessage, []byte(err))
	c.Ws.Close()
}

func (c *Client) SendError(err string){
	c.Ws.WriteMessage(websocket.TextMessage, []byte(err))
}


//
func (c *Client) RemoveAfter(duration int, hub *Hub) {
	d := time.Duration(duration)
	time.Sleep(d * time.Second)
	hub.RemoveClient <- c
	if hub.Clients[c] {
		c.Ws.Close()
	}
}
