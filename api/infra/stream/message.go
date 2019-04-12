package stream

type SocketMessage struct {
	Token string `json:"token"`
	Event string `json:"event"`

	//Get Offer
	Page    int    `json:"page"`
	Size    int    `json:"size"`
	Sortkey string `json:"sort_key"`
}

type HubMessage struct {
	Data *SocketMessage
	Client  *Client
	Hub *Hub
}

type SocketErrorMessage struct{
	Error string
}