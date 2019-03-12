package user

import (
	"auction/pkg/entity"
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
func (r *MongoRepository) Find(id entity.ID) (*User, error) {
	result := User{}
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("users")
	err := coll.Find(bson.M{"_id": id}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}

//Store : Insert an offer
func (r *MongoRepository) Save(user *User) (entity.ID, error) {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("users")
	err := coll.Insert(user)
	if err != nil {
		return entity.ID(0), err
	}
	return user.ID, nil
}

//FindByKey
func (r *MongoRepository) FindByKey(key string, val interface{}) ([]*User, error) {
	var result []*User
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("users")
	err := coll.Find(bson.M{key: val}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}
