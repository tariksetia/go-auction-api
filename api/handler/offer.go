package handler

import (
	"auction/api/infra/hub"
	e "auction/pkg/entity"
	"auction/pkg/offer"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{}

func createOffer(hub *hub.Hub, service offer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ofr *e.Offer
		usr := r.Context().Value("me").(*e.User)
		errorMessage := "Error Creating Ofr"
		err := json.NewDecoder(r.Body).Decode(&ofr)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Creating Offer"))
			return
		}

		// check if offer data is valid else return error
		if !ofr.Validate() {
			log.Println("Invalid Offer Data")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid Offer Data"))
			return
		}
		ofr.CreatedBy = usr.Username
		ofr.Id, err = service.Save(ofr)

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if err := json.NewEncoder(w).Encode(ofr); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
		w.Header().Set("Content-Type", "application/json")
		hub.Broadcast <- []byte("OfferCreated")
		w.WriteHeader(http.StatusCreated)

	})
}

func getOffer(service offer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		errorMessage := "Error reading Offers"
		var ofrs []*e.Offer
		page, err := strconv.Atoi(r.FormValue("page"))
		if err != nil {
			page = 0
		}
		size, err := strconv.Atoi(r.FormValue("size"))
		if err != nil {
			size = 10
		}
		sortKey := r.FormValue("sortKey")

		ofrs, err = service.Query(page, size, sortKey)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusAccepted)
		if err := json.NewEncoder(w).Encode(ofrs); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//CreateUserHandlers Maps routes to http handlers
func CreateOfferHandlers(hub *hub.Hub, r *mux.Router, n negroni.Negroni, service offer.UseCase) {
	r.Handle("/v1/offer", n.With(
		negroni.Wrap(createOffer(hub, service)),
	)).Methods("POST", "OPTIONS").Name("CreateOffer")

	r.Handle("/v1/offer", n.With(
		negroni.Wrap(getOffer(service)),
	)).Methods("GET", "OPTIONS").Name("GetOffers")
}
