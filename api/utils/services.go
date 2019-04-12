package utils

import (
	"auction/pkg/bid"
	"auction/pkg/offer"
	"auction/pkg/user"
)

type Services struct {
	User  user.Service
	Offer offer.Service
	Bid   bid.Service
}
