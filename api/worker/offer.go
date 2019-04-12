package worker

import (
	"auction/api/utils"
	e "auction/pkg/entity"
	"log"
)

func consumeBids(msg *e.BidChannelMessage, services *utils.Services) {
	/*
		Process HTTP POSTED bids

		_bid: The bid placed by user
		bidSerice: acces to bid domain logic
		offerService: access to offer domain logic
		me: User infro decoded from JWT by the middleware, and accessible to handler
	*/
	bid := msg.BidEntity
	offer := msg.OfferEntity

	bidService := services.Bid
	offerService := services.Offer

	//Check if current bid_price > old bid price
	//We save the bid and update the offers last bid price if it is lesser than the current bid_price
	if offer.BidPrice >= bid.BidPrice {
		bid.Accepted = false
	} else {
		bid.Accepted = true
	}

	var err error
	if bid.Accepted {
		offer, err = offerService.Update(bid.OfferID, "bidprice", bid.BidPrice)
		if err != nil {
			log.Println("Error Placing Bid")
			return
		}
	}

	//save the bid
	bid.Id, err = bidService.Save(bid)
	if err != nil {
		log.Println("Error Placing Bid")
		return
	}
}
