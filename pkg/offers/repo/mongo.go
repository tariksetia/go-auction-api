package repo

import (
	"auction/pkg/entity"
	"auction/pkg/offers"
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
func NewMongoRepository(p *mgosession.Pool, db string) offers.Repository {
	return &MongoRepository{
		pool: p,
		db:   db,
	}
}

//Get a offer by ID
func (r *MongoRepository) Find(id entity.ID) (*offers.Offer, error) {
	result := offers.Offer{}
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("offers")
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

//Insert an offer
func (r *MongoRepository) Store(b *offers.Offer) (entity.ID, error) {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("offers")
	err := coll.Insert(b)
	if err != nil {
		return entity.ID(0), err
	}
	return b.ID, nil
}

//query offers
func (r *MongoRepository) Query(page int, size int, sortkey string) ([]*offers.Offer, error) {

	if size == 0 {
		size = 10
	}

	if sortkey == "" {
		sortkey = "go_live"
	}

	var res []*offers.Offer
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("offers")
	err := coll.Find(nil).Sort(sortkey).Limit(size).Skip(page).All(&res)
	switch err {
	case nil:
		return res, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}
