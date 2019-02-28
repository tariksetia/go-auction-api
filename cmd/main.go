package main

import (
	"auction/pkg/offers"
	"auction/pkg/offers/repo"
	"fmt"
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
	"log"
	"time"
)

func main() {

	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	mPool := mgosession.NewPool(nil, session, 5)
	defer mPool.Close()
	db := "auction"
	offerRepo := repo.NewMongoRepository(mPool, db)
	offerService := offers.NewService(offerRepo)
	/*
		type Offer struct {
				ID       entity.ID `json:"id"`
				BidPrice float64   `json:"bid_price"`
				GoLive   time.Time `json:"go_live"`
				Lifetime int64     `json:"lifetime"`
				PhotoUrl string    `json:"photo_url"`
				Title    string    `json:"title"`
			}
	*/
	testoffer := offers.Offer{
		BidPrice: 12.34,
		GoLive:   time.Now(),
		Lifetime: 23,
		PhotoUrl: "/photo/url",
		Title:    "LoremIpsum",
	}
	id, err := offerService.Store(&testoffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(id))

}
