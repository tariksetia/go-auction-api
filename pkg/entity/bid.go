package entity

type Bid struct {
	ID       int64   `json:"id"`
	BidPrice float64 `json:"bid_price"`
	ClientId string  `json:"client_id"`
	OfferId  string  `json:"offer_id"`
}
