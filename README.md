## Go Auction API

### MISC Routes
- GET /ping: confirms that server is up and running
- GET /index: html page for testing websocket events


### API Routes
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
    }```

    Response

    ```json
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkIjoxNTU0MTk5NjA2LCJ1c2VySUQiOiI1Y2EzMzQzMjc2ODg5NzMxZjhhMjhmYzEifQ.1OPAlenPQu0hmLAcMrXYeKyNZK0WxAulIUhuNZgoVFA"
    }
    '''
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
    '''

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