package offers

import (
	"auction/pkg/entity"
	"time"
)

//Offer Entity Structure
type Offer struct {
	ID       entity.ID `json:"id"`
	BidPrice float64   `json:"bid_price"`
	GoLive   time.Time `json:"go_live"`
	Lifetime int64     `json:"lifetime"`
	PhotoUrl string    `json:"photo_url"`
	Title    string    `json:"title"`
}
