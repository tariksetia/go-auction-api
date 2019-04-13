package stream

import (
	"auction/api/utils"
	//e "auction/pkg/entity"
	"encoding/json"
	"log"
)

type Hub struct {
	Clients  map[*Client]bool
	Services *utils.Services

	AddClient    chan *Client
	RemoveClient chan *Client

	Incoming     chan *HubMessage
	Authenticate chan *HubMessage
	GetOffers    chan *HubMessage

	Broadcast     chan []byte
	UnicastJSON   chan *UnicastMessage
	BroadcastJSON chan *SocketOutGoingMessage
	started       bool
}

var hub = Hub{
	Clients: make(map[*Client]bool),

	AddClient:    make(chan *Client),
	RemoveClient: make(chan *Client),

	Incoming:     make(chan *HubMessage),
	Authenticate: make(chan *HubMessage),
	GetOffers:    make(chan *HubMessage),

	Broadcast:     make(chan []byte),
	UnicastJSON:   make(chan *UnicastMessage),
	BroadcastJSON: make(chan *SocketOutGoingMessage),
}

func (hub *Hub) start() {
	hub.started = true
	for {
		select {
		case msg := <-hub.Incoming:
			go processIncomingMessage(msg)
		case msg := <-hub.Authenticate:
			go processSocketAuthentication(msg)
		case msg := <-hub.GetOffers:
			go processSocketGetOffers(msg)
		case conn := <-hub.AddClient:
			hub.Clients[conn] = true

		case conn := <-hub.RemoveClient:
			if _, ok := hub.Clients[conn]; ok {
				delete(hub.Clients, conn)
				//close(conn.Send)
			}
		case message := <-hub.Broadcast:
			for conn := range hub.Clients {
				conn.Send <- message
				log.Println(message)
			}
		case message := <-hub.BroadcastJSON:
			for conn := range hub.Clients {
				conn.Ws.WriteJSON(message)
			}
		case message := <-hub.UnicastJSON:
			sendPrivateMessate(message)
		}
	}
}

func GetHub(services *utils.Services) *Hub {
	if hub.started {
		return &hub
	}
	hub.Services = services
	go hub.start()
	return &hub
}

func sendPrivateMessate(msg *UnicastMessage) {
	//Send private message to a socket connection
	//it could be an error or message or contains query data
	msg.Client.Ws.WriteJSON(msg.Message)

	//If kill == true, mostly in cases of errors, kill/close the websocket
	if msg.Message.Kill {
		delete(hub.Clients, msg.Client)
		close(msg.Client.Send)
	}
}
func processIncomingMessage(msg *HubMessage) {
	event := string(msg.Data.Event)
	if event == "authenticate" {
		msg.Hub.Authenticate <- msg
		return
	}
	if event == "get_offers" {
		msg.Hub.GetOffers <- msg
		return
	}

	//If Websocket is Unauthenticated, close Socket
	//else return error Message over socket
	outMsg := SocketOutGoingMessage{
		Error: "Invalid Event Type",
	}
	uniMsg := UnicastMessage{
		Client:  msg.Client,
		Message: &outMsg,
	}
	if !msg.Hub.Clients[msg.Client] {
		outMsg.Kill = true
	}
	msg.Hub.UnicastJSON <- &uniMsg
}

func processSocketAuthentication(msg *HubMessage) {
	token := msg.Data.Token

	delta, err := utils.ValidateAuthToken(token)

	outMsg := SocketOutGoingMessage{}
	uniMsg := UnicastMessage{
		Client:  msg.Client,
		Message: &outMsg,
	}
	if err != nil {
		//token validation error
		//remove the client if exist
		//close the connection with error
		msg.Hub.RemoveClient <- msg.Client

		outMsg.Error = err.Error()
		outMsg.Kill = false

		msg.Hub.UnicastJSON <- &uniMsg
		return
	}
	if msg.Hub.Clients[msg.Client] == true {
		/*
			at this point token is valid but this is
			redundant request as client already exist in Hub.Clients.
			In other words, client is already authenticated
		*/
		outMsg.Message = "Already Authenticated"
		msg.Hub.UnicastJSON <- &uniMsg
		return
	}

	//Else Add the client to Hub.Client
	msg.Hub.AddClient <- msg.Client

	outMsg.Message = "User Authenticated"
	msg.Hub.UnicastJSON <- &uniMsg

	//Start a goroutine that closes the socket and remove
	// client form Hub.Client once token expires
	go msg.Client.RemoveAfter(delta, msg.Hub)

}

func processSocketGetOffers(msg *HubMessage) {

	page := msg.Data.Page
	size := msg.Data.Size
	sortKey := msg.Data.Sortkey
	client := msg.Client
	outMsg := SocketOutGoingMessage{}
	uniMsg := UnicastMessage{
		Message: &outMsg,
		Client:  msg.Client,
	}

	if !msg.Hub.Clients[client] {
		//This means client is not authenticated
		//Kill the socket as it is unauthenticated
		outMsg.Kill = true
		outMsg.Error = "Websocket Connection is unauthenticated"

		msg.Hub.UnicastJSON <- &uniMsg
		return
	}
	if size == 0 {
		size = 10
	}
	if sortKey == "" {
		sortKey = "golive"
	}

	ofrs, err := msg.Hub.Services.Offer.Query(page, size, sortKey)
	if err != nil {
		outMsg.Error = err.Error()
		msg.Hub.UnicastJSON <- &uniMsg
	}

	data, _ := json.Marshal(ofrs)
	outMsg.Data = string(data)

	msg.Hub.UnicastJSON <- &uniMsg
	return
}
