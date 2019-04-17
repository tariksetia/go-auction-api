package auc_test_test

import (
	"auction/api/config"
	"auction/api/handler"
	"auction/api/middleware"
	"auction/api/stream"
	"auction/api/utils"
	"auction/api/worker"
	"auction/mongo"
	"auction/pkg/bid"
	"auction/pkg/offer"
	"auction/pkg/user"
	testutils "auction/utils"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/juju/mgosession"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/mgo.v2"
	"testing"
)

var mPool *mgosession.Pool
var session *mgo.Session
var r *mux.Router
var services utils.Services
var hub *stream.Hub
var broker *worker.Broker
var authMiddleware *negroni.Negroni
var apiMiddleware *negroni.Negroni

func TestAucTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AucTest Suite")
}

var _ = BeforeSuite(func() {
	testutils.MongoStart()

	cfg := config.AppConfig{
		DBhost:         "0.0.0.0",
		DBname:         "auction",
		DBport:         27017,
		ConnectionPool: 1,
		AppSecret:      "SAMSUNG-H45-A-100-%-F0LD1NG-PH0N3",
		AppServerPort:  8000,
	}

	mPool, session = mongo.GetMongoPool(
		cfg.GetDatabaseHostname(),
		cfg.GetDatabasePort(),
		cfg.GetConnectionPool(),
	)

	//Create MUX router
	r = mux.NewRouter()

	userRepo := user.CreateMongoRepo(mPool, cfg.GetDatabaseName())
	offerRepo := offer.CreateMongoRepository(mPool, cfg.GetDatabaseName())
	bidRepo := bid.CreateMongoRepository(mPool, cfg.GetDatabaseName())

	userService := user.NewService(userRepo)
	offerService := offer.NewService(offerRepo)
	bidService := bid.NewService(bidRepo)

	services = utils.Services{
		User:  *userService,
		Offer: *offerService,
		Bid:   *bidService,
	}

	//Create The Hub
	hub = stream.GetHub(&services)

	//Create the Message Broker
	broker = worker.GetOrCreateBroker(&services)

	//Middleware for signup and login
	authMiddleware = negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	//Middleware for all other routes that require authentication
	apiMiddleware = negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.JwtMiddleware(cfg.AppSecret)),
		negroni.HandlerFunc(middleware.LoginMiddleware(userService)),
		negroni.NewLogger(),
	)

	//create Handlers for different domain services
	handler.CreateUserHandlers(r, *authMiddleware, userService)
	handler.CreateOfferHandlers(hub, r, *apiMiddleware, offerService)
	handler.CreateBidHandlers(r, *apiMiddleware, broker, bidService, offerService)
	handler.CreateStreamHandler(r, *authMiddleware, hub, &services)

})

var _ = AfterSuite(func() {
	testutils.MongoKill()
	testutils.MongoRemove()
	mPool.Close()
	session.Close()
})
