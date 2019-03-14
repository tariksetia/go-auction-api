package entity

import (
	"time"
)

//Offer Entity Structure
type Bid struct {
	Id        ID        `json:"id" bson:"_id,omitempty"`
	BidPrice  float64   `json:"bid_price"`
	Username  string    `json:"username"`
	OfferID   ID        `json:"offer_id"`
	Timestamp time.Time `json:"timestamp"`
	Accepted  bool      `json:"accepted"`
}

//Validate Validate
func (bid *Bid) Validate() bool {
	if bid.BidPrice <= 0 || bid.OfferID == "" {
		return false
	}
	bid.Timestamp = time.Now()
	return true
}
