package bid

import (
	e "auction/pkg/entity"
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

//MongoRepository mongodb repo
type MongoRepository struct {
	pool *mgosession.Pool
	db   string
}

//NewMongoRepository create new repository
func CreateMongoRepository(p *mgosession.Pool, db string) Repository {
	return &MongoRepository{
		pool: p,
		db:   db,
	}
}

//Find : Get a Bid by ID
func (r *MongoRepository) Find(id e.ID) (*e.Bid, error) {
	result := e.Bid{}
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("bids")
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

//Store : Insert an Bid
func (r *MongoRepository) Save(b *e.Bid) (e.ID, error) {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("bids")
	err := coll.Insert(b)
	if err != nil {
		return e.ID(0), err
	}
	return b.Id, nil
}

//FindByKey
func (r *MongoRepository) FindByKey(key string, val interface{}, page int, size int) ([]*e.Bid, error) {

	if size == 0 {
		size = 10
	}

	var res []*e.Bid
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("bids")
	err := coll.Find(bson.M{key: val}).Limit(size).Skip(page).All(&res)
	switch err {
	case nil:
		return res, nil
	case mgo.ErrNotFound:
		return nil, e.ErrNotFound
	default:
		return nil, err
	}
}

//Update
func (r *MongoRepository) Update(id e.ID, key string, val interface{}) (*e.Bid, error) {
	result := e.Bid{}
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("bids")
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{key: val}},
		ReturnNew: true,
	}
	_, err := coll.Find(bson.M{"_id": id}).Apply(change, &result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, e.ErrNotFound
	default:
		return nil, err
	}
}
