package main

import (
	"auction/pkg/offer"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"auction/api/handler"
	"auction/pkg/user"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	r := mux.NewRouter()

	mPool := mgosession.NewPool(nil, session, 50)
	defer mPool.Close()

	userRepo := user.CreateMongoRepo(mPool, "auction")
	userService := user.NewService(userRepo)
	offerRepo := offer.CreateMongoRepository(mPool, "auction")
	offerService := offer.NewService(offerRepo)
	//handlers
	n := negroni.New(
		negroni.NewLogger(),
	)
	handler.CreateUserHandlers(r, *n, userService)
	handler.CreateOfferHandlers(r, *n, offerService)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(8000),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
