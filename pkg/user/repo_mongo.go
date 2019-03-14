package user

import (
	e "auction/pkg/entity"
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

type MongoRepository struct {
	pool *mgosession.Pool
	db   string
}

//NewMongoRepository create new repository
func CreateMongoRepo(p *mgosession.Pool, db string) Repository {
	return &MongoRepository{
		pool: p,
		db:   db,
	}
}

//Find : Get a offer by ID
func (r *MongoRepository) Find(id e.ID) (*e.User, error) {
	result := e.User{}
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("users")
	err := coll.Find(bson.M{"_id": id}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, e.ErrNotFound
	default:
		return nil, err
	}
}

//Store : Insert an offer
func (r *MongoRepository) Save(user *e.User) (e.ID, error) {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("users")
	err := coll.Insert(user)
	if err != nil {
		return e.ID(0), err
	}
	return user.Id, nil
}

//FindByKey
func (r *MongoRepository) FindByKey(key string, val interface{}) ([]*e.User, error) {
	var result []*e.User
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("users")
	err := coll.Find(bson.M{key: val}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, e.ErrNotFound
	default:
		return nil, err
	}
}
