package main

import (
	"auction/api/config"
	"auction/api/handler"
	h "auction/api/infra/hub"
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

	//Create The Hub
	hub := h.GetHub()

	//Create MUX router
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

	//Middleware for all other routes that require authentication
	apiMiddleware := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.JwtMiddleware(cfg)),
		negroni.HandlerFunc(middleware.LoginMiddleware(userService)),
		negroni.NewLogger(),
	)

	//create Handlers for different domain services
	handler.CreateUserHandlers(r, *authMiddleware, userService)
	handler.CreateOfferHandlers(hub, r, *apiMiddleware, offerService)
	handler.CreateBidHandlers(r, *apiMiddleware, bidService, offerService)
	handler.CreateStreamHandler(r, *authMiddleware, hub)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + cfg.GetAppServerPort(),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Server is UP!!!!")
}
