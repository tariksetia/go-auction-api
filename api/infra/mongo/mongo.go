package mongo

import (
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
	"log"
)

func GetMongoPool(numConnections int) (*mgosession.Pool, *mgo.Session) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()
	mPool := mgosession.NewPool(nil, session, numConnections)
	return mPool, session
}
