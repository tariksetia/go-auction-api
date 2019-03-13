package main

import (
	"auction/api/handler"
	"auction/api/middleware"
	"auction/pkg/bid"
	"auction/pkg/offer"
	"auction/pkg/user"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	r := mux.NewRouter()

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("this_is_the_end_hold_your_breath_and_count_to_10"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	mPool := mgosession.NewPool(nil, session, 50)
	defer mPool.Close()

	userRepo := user.CreateMongoRepo(mPool, "auction")
	userService := user.NewService(userRepo)
	offerRepo := offer.CreateMongoRepository(mPool, "auction")
	offerService := offer.NewService(offerRepo)
	bidRepo := bid.CreateMongoRepository(mPool, "auction")
	bidService := bid.NewService(bidRepo)
	//handlers
	authMiddleware := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)
	apiMiddleware := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.HandlerFunc(middleware.Login),
		negroni.NewLogger(),
	)
	handler.CreateUserHandlers(r, *authMiddleware, userService)
	handler.CreateOfferHandlers(r, *apiMiddleware, offerService)
	handler.CreateBidHandlers(r, *apiMiddleware, bidService, offerService)

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
