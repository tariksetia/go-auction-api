package mongo

import (
	"fmt"
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
	"log"
)

//GetMongoPool Create reference mongodb session and connection pool
func GetMongoPool(host string, port string, numConnections int) (*mgosession.Pool, *mgo.Session) {
	connectionString := fmt.Sprintf("mongodb://%s:%s", host, port)
	session, err := mgo.Dial(connectionString)
	if err != nil {
		log.Fatal("Cannot Connect to MongdoDB")

	}
	mPool := mgosession.NewPool(nil, session, numConnections)
	return mPool, session
}
