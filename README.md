# Go Auction API


## Setup

docker-compose up

## MISC Routes
- GET /ping: confirms that server is up and running
- GET /index: html page for testing websocket events


## API Routes
- POST /v1/signup

    Request
    ```json
    {
        "username":"Azerbaijan",
        "password":"baku"
    }
    ```

- POST /v1/login
    
    Request
    ```json
    {
        "username":"Azerbaijan",
        "password":"baku"
    }
    ```

    Response

    ```json
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkIjoxNTU0MTk5NjA2LCJ1c2VySUQiOiI1Y2EzMzQzMjc2ODg5NzMxZjhhMjhmYzEifQ.1OPAlenPQu0hmLAcMrXYeKyNZK0WxAulIUhuNZgoVFA"
    }
    ```
- POST /v1/offer

    Request
    ```json
    {
        "bid_price": 12.34, //(Required)
        "lifetime": 0,      //(Optional)
        "photo_url": "",    //(Optional)
        "title": "Swiss Shales", //(Required)
    }
    ```
    Response
    ```json
    {
        "id": "5ca3414276889736ba1d7efd",
        "bid_price": 12.34,
        "go_live": "2019-04-02T16:32:26.596273+05:30",
        "lifetime": 0,
        "photo_url": "",
        "title": "THis is a test",
        "created_by": "tariksetia",
        "sold": false
    }
    ```

- GET /v1/offer or /v1/offeroffer?size=10&page=0&sortKey=golive

- POST /v1/bid
    
    Request
    ```json
    {
        "bid_price":6832427,
        "offer_id":"5ca34ac07688973b556056ce"
    }
    ```

    Respone
    ```json
    {
        "id": "5ca34c2f7688973f34f8dee1",
        "bid_price": 6832427,
        "username": "tariksetia",
        "offer_id": "5ca34ac07688973b556056ce",
        "timestamp": "2019-04-02T17:19:03.799762+05:30",
        "accepted": false
    }
    ```

- PUT /vi/bid/:bidid
    
    RESPONSE
    ```json
    {
        "id": "5ca34c2f7688973f34f8dee1",
        "bid_price": 6832427,
        "username": "tariksetia",
        "offer_id": "5ca34ac07688973b556056ce",
        "timestamp": "2019-04-02T17:19:03.799+05:30",
        "accepted": true
    }
    ```


## Websocket
- The incoming message from websocket is of type ```SocketMessage```
- ```SocketMessage``` consist of ```Event``` and ```Token``` fileds along with other fields depengding on the ``Event`` type
- ```Token``` Each socket message must contain valid token, else socket stream is closed
- ```Event``` contains the event type a ```String```, that tell what processing is to be done
- There are only two implemented websocket events ```authenticate``` and ```get_offer```
- The very first event type when socket connects must be ```authenticate```, socket connection will be closed if event type is something else
- Successful authentication adds socket to global websocket client list and each time an offer is created, ```offerCreated``` message is broadcasted to every client
- ```Hub`` is centeral message reciever and dispatcher, consisting of various channels, list of authenticated client and access to all domain services
- It consist of following channels
    - ```Clients  map[*Client]bool``` : Authenticated clients map
	- ```Services *utils.Services```  : Reference to domain services
    
	- ```AddClient    chan *Client``` 
	- ```RemoveClient chan *Client```

	- ```Incoming     chan *HubMessage```
	- ```Authenticate chan *HubMessage```
	- ```GetOffers    chan *HubMessage```

	- ```Broadcast     chan []byte```
	- ```UnicastJSON   chan *UnicastMessage```
	- ```BroadcastJSON chan *SocketOutGoingMessage```
	- ```started       bool``` : Falg to tell whether hub is running or not

- All the incoming messages Fans in at ```Incoming``` channel and based on the event type faned-out to different channels

## Broker
- ```Broker``` package mimics message queues using channels, It is ok for the sake of this POC, but try to use standard bokers such as ActiveMQ, RabbitMq or Kafka in production. NATS dont provide message persistence
- Every ```Bid``` is queued once succefully validated
- Later repective queue workers (goroutines) process the message.
- Every ```Bid``` is proccessed sequentially and highest bid is updated in ```Offers```

## Benchmarking
```
go get -u github.com/rakyll/hey


./hey -m POST -n 100000 -c 1000 \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer <TOKEN>" \
        -D ~/Lab/data.json \
        http://localhost:8000/v1/
```