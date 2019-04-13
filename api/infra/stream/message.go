package stream

//The incoming Message from websocket client
//One Struct for All event request
//Send Fields that relates to event, extra fields will be ignored
type SocketMessage struct {
	Token string `json:"token"`
	Event string `json:"event"`

	//Get Offer
	Page    int    `json:"page"`
	Size    int    `json:"size"`
	Sortkey string `json:"sort_key"`
}

//HubMessage: A workaround data structure for passign around between channels
type HubMessage struct {
	Data   *SocketMessage //Pointer to Incoming Message
	Client *Client        //Pointer to Client for sending message
	Hub    *Hub           //Pointer to Hub for accessing different channel
}

//SocketOutGoingMessage: Data Struct used by WriteJson for Unicast and BroadCast
type SocketOutGoingMessage struct {
	Data    interface{} `json:"data"` // Stringified query data TODO:: Fix it, to send JSON
	Error   interface{} `json:"error"`
	Code    interface{} `json:"code"`
	Message interface{} `json:"message"`
	Kill    bool        // If kill is true, websocket will be closed after delivering message
}

//UnicastMessage: Contains reference to cliend and message to be send
type UnicastMessage struct {
	Client  *Client
	Message *SocketOutGoingMessage
}
