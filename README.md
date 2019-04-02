#Go Auction API

###MISC Routes
- GET /ping: confirms that server is up and running
- GET /index: html page for testing websocket events


## API Routes
- POST /v1/signup
    REQUEST = {
        "username":"Azerbaijan",
        "password":"baku"
    }

- POST /v1/login
    REQUEST = {
        "username":"Azerbaijan",
        "password":"baku"
    }

    RESPONSE = {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkIjoxNTU0MTk5NjA2LCJ1c2VySUQiOiI1Y2EzMzQzMjc2ODg5NzMxZjhhMjhmYzEifQ.1OPAlenPQu0hmLAcMrXYeKyNZK0WxAulIUhuNZgoVFA"
    }

- POST /v1/offer
    REQUEST = {
        "bid_price": 12.34, (Required)
        "lifetime": 0,      (Optional)
        "photo_url": "",    (Optional)
        "title": "Swiss Shales", (Required)
    }
    RESPONSE = {
        "id": "5ca3414276889736ba1d7efd",
        "bid_price": 12.34,
        "go_live": "2019-04-02T16:32:26.596273+05:30",
        "lifetime": 0,
        "photo_url": "",
        "title": "THis is a test",
        "created_by": "tariksetia",
        "sold": false
    }

- GET /v1/offer or



  