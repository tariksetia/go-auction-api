package handler

import (
	"auction/api/worker"
	"auction/pkg/bid"
	e "auction/pkg/entity"
	"auction/pkg/offer"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"log"
	"net/http"
)

func placeBid(bidService bid.UseCase, offerService offer.UseCase, broker *worker.Broker) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var _bid *e.Bid
		//errorMessage := "Error Placing Bid"
		usr := r.Context().Value("me").(*e.User)
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
		offer, err := offerService.Find(_bid.OfferID)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Placing Bid, Could not find any offer"))
			return
		}

		_bid.Username = usr.Username
		// Add it to channel
		msg := e.BidChannelMessage{
			BidEntity:   _bid,
			OfferEntity: offer,
		}

		broker.BidQueue <- &msg

		w.WriteHeader(http.StatusCreated)

	})
}

func acceptBid(bidService bid.UseCase, offerService offer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		bidID := e.StringToID(id)
		errorMessage := "Error Accepting Bid"

		// update Bid if valid and greater than previous
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
//Broker required as we are queueing all the bids and then processing them
//TODO: Move to a persistent message broker, for this POC go channel is sufficient
func CreateBidHandlers(r *mux.Router, n negroni.Negroni, broker *worker.Broker, bidService bid.UseCase, offerService offer.UseCase) {
	r.Handle("/v1/bids", n.With(
		negroni.Wrap(placeBid(bidService, offerService, broker)),
	)).Methods("POST", "OPTIONS").Name("placeBid")

	r.Handle("/v1/bids/{id}", n.With(
		negroni.Wrap(acceptBid(bidService, offerService)),
	)).Methods("PUT", "OPTIONS").Name("acceptBid")

}
