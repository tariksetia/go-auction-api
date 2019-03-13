package main

import (
	"auction/api/config"
	"auction/api/handler"
	"auction/api/infra/mongo"
	"auction/api/middleware"
	"auction/pkg/bid"
	"auction/pkg/offer"
	"auction/pkg/user"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	//Get Application Config
	cfg := config.GetAppConfig()

	//Connect to MongoDB
	mPool, session := mongo.GetMongoPool(
		cfg.GetDatabaseHostname(),
		cfg.GetDatabasePort(),
		cfg.GetConnectionPool(),
	)
	defer session.Close()
	defer mPool.Close()

	r := mux.NewRouter()
	userRepo := user.CreateMongoRepo(mPool, cfg.GetDatabaseName())
	offerRepo := offer.CreateMongoRepository(mPool, cfg.GetDatabaseName())
	bidRepo := bid.CreateMongoRepository(mPool, cfg.GetDatabaseName())

	userService := user.NewService(userRepo)
	offerService := offer.NewService(offerRepo)
	bidService := bid.NewService(bidRepo)

	//Middleware for signup and login
	authMiddleware := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)
	apiMiddleware := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.JwtMiddleware(cfg)),
		negroni.HandlerFunc(middleware.LoginMiddleware(userService)),
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
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
