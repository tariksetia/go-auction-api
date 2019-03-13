package handler

import (
	"auction/pkg/bid"
	"auction/pkg/entity"
	"auction/pkg/offer"
	"auction/pkg/user"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"log"
	"net/http"
)

func placeBid(bidService bid.UseCase, offerService offer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var _bid *bid.Bid
		errorMessage := "Error Placing Bid"
		usr := r.Context().Value("me").(*user.User)
		err := json.NewDecoder(r.Body).Decode(&_bid)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Placing Bid"))
			return
		}

		// check if offer data is valid else return error
		if !_bid.Validate() {
			log.Println("Error Placing Bid")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Placing Bid"))
			return
		}

		//get the offer by ID
		ofr, err := offerService.Find(_bid.OfferID)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Placing Bid"))
			return
		}

		//check if current bid_price > old bid price
		if ofr.BidPrice >= _bid.BidPrice {
			log.Println("Error Placing Bid. BidPrice is lesser than previous value")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Placing Bid. BidPrice is lesser than previous value"))
			return
		}

		//update the bid_price in offer
		ofr, err = offerService.Update(_bid.OfferID, "bidprice", _bid.BidPrice)
		if err != nil {
			log.Println("Error Placing Bid")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Placing Bid"))
			return
		}
		//save the bid
		_bid.Username = usr.Username
		_bid.ID, err = bidService.Save(_bid)
		if err != nil {
			log.Println("Error Placing Bid")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Placing Bid"))
			return
		}

		if err := json.NewEncoder(w).Encode(_bid); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

	})
}

func acceptBid(bidService bid.UseCase, offerService offer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		bidID := entity.StringToID(id)
		errorMessage := "Error Accepting Bid"

		// update Bid
		_bid, err := bidService.Update(bidID, "accepted", true)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Placing Bid"))
			return
		}
		//Update Offer
		_, err = offerService.Update(_bid.OfferID, "sold", true)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Placing Bid"))
			return
		}

		if err := json.NewEncoder(w).Encode(_bid); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

	})
}

//CreateUserHandlers Maps routes to http handlers
func CreateBidHandlers(r *mux.Router, n negroni.Negroni, bidService bid.UseCase, offerService offer.UseCase) {
	r.Handle("/v1/bids", n.With(
		negroni.Wrap(placeBid(bidService, offerService)),
	)).Methods("POST", "OPTIONS").Name("placeBid")

	r.Handle("/v1/bids/{id}", n.With(
		negroni.Wrap(acceptBid(bidService, offerService)),
	)).Methods("PUT", "OPTIONS").Name("acceptBid")

}
