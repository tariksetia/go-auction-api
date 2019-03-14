package entity

import (
	"time"
)

//Offer Entity Structure
type Offer struct {
	Id        ID        `json:"id" bson:"_id,omitempty"`
	BidPrice  float64   `json:"bid_price"`
	GoLive    time.Time `json:"go_live"`
	Lifetime  int64     `json:"lifetime"`
	PhotoURL  string    `json:"photo_url"`
	Title     string    `json:"title"`
	CreatedBy string    `json:"created_by"`
	Sold      bool      `json:"sold"`
}

//Validate : Check if data in offer is valid or not, make adjustment whereever possible
func (offer *Offer) Validate() bool {
	if offer.BidPrice == 0 || offer.Title == "" || offer.Lifetime < 0 {
		return false
	}
	if offer.GoLive.Before(time.Now()) {
		offer.GoLive = time.Now()
	}
	return true
}
