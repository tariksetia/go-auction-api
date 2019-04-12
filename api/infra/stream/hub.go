package stream

import(
	"auction/api/utils"
	//e "auction/pkg/entity"
	"fmt"
	"encoding/json"
)
type Hub struct {
	Clients      map[*Client]bool
	Services	 *utils.Services
	
	AddClient    chan *Client
	RemoveClient chan *Client

	Incoming     chan *HubMessage
	Authenticate chan *HubMessage
	GetOffers	 chan *HubMessage

	Broadcast    chan []byte
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
		case msg := <-hub.Incoming:
			fmt.Println("Processing Incoming Message")
			processIncomingMessage(msg)
		case msg := <-hub.Authenticate:
			processSocketAuthentication(msg)
		case conn := <-hub.AddClient:
			hub.Clients[conn] = true
		case conn := <-hub.RemoveClient:
			if _, ok := hub.Clients[conn]; ok {
				delete(hub.Clients, conn)
				close(conn.Send)
			}
		case message := <-hub.Broadcast:
			for conn := range hub.Clients {
				conn.Send <- message
			}
		
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

func processIncomingMessage(msg *HubMessage) {
	switch msg.Data.Event {
	case "authenticate":{
		msg.Hub.Authenticate <- msg
	}
	case "get_offers":{
		msg.Hub.GetOffers <- msg
	}
}
}

func processSocketAuthentication(msg *HubMessage){
	token :=msg.Data.Token
	
	delta, err := utils.ValidateAuthToken(token)
	if err != nil {
		msg.Client.CloseWithError(err.Error())
	}

	msg.Hub.AddClient <- msg.Client
	go msg.Client.RemoveAfter(delta, msg.Hub)

}

func processSocketGetOffers(msg *HubMessage){
	
	page := msg.Data.Page
	size := msg.Data.Size
	sortKey := msg.Data.Sortkey
	client := msg.Client


	if  !msg.Hub.Clients[client]{
		//This means client is not authenticated
		client.CloseWithError("Unauthenticated Socket")
	}
	if size == 0 {
		size = 10
	}
	if sortKey == ""{
		sortKey = "golive"
	}


	ofrs, err := msg.Hub.Services.Offer.Query(page, size, sortKey)
	if err != nil {
		client.SendError(err.Error())
	}
	data,_ := json.Marshal(ofrs)
	fmt.Println(data)
}