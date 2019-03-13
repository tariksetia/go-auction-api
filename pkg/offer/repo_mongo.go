package offer

import (
	"auction/pkg/entity"
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

//Find : Get a offer by ID
func (r *MongoRepository) Find(id entity.ID) (*Offer, error) {
	result := Offer{}
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

//Store : Insert an offer
func (r *MongoRepository) Save(b *Offer) (entity.ID, error) {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("offers")
	err := coll.Insert(b)
	if err != nil {
		return entity.ID(0), err
	}
	return b.ID, nil
}

//Query :query offers
func (r *MongoRepository) Query(page int, size int, sortkey string) ([]*Offer, error) {

	if size == 0 {
		size = 10
	}

	if sortkey == "" {
		sortkey = "go_live"
	}

	var res []*Offer
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

//FindByKey
func (r *MongoRepository) FindByKey(key string, val interface{}, page int, size int) ([]*Offer, error) {

	if size == 0 {
		size = 10
	}

	var res []*Offer
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("offers")
	err := coll.Find(bson.M{key: val}).Limit(size).Skip(page).All(&res)
	switch err {
	case nil:
		return res, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}

//Update
func (r *MongoRepository) Update(id entity.ID, key string, val interface{}) (*Offer, error) {
	result := Offer{}
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("offers")
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{key: val}},
		ReturnNew: true,
	}
	_, err := coll.Find(bson.M{"_id": id}).Apply(change, &result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}
