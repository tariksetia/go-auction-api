package repo

import (
	"auction/pkg/auth"
	"auction/pkg/entity"
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

//MongoRepository mongodb repo
type AuthMongoRepository struct {
	pool *mgosession.Pool
	db   string
}

//NewMongoRepository create new repository
func CreateAuthMongoRepo(p *mgosession.Pool, db string) auth.Repository {
	return &AuthMongoRepository{
		pool: p,
		db:   db,
	}
}

//Find : Get a User by ID
func (r *AuthMongoRepository) FindByUserid(id entity.ID) (*auth.User, error) {
	result := auth.User{}
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

//FindByUsername : Get a User by ID
func (r *AuthMongoRepository) FindByUsername(username string) (*auth.User, error) {
	result := auth.User{}
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("users")
	err := coll.Find(bson.M{"username": username}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}

//FindByToken : Get a User info using security token
func (r *AuthMongoRepository) FindByToken(username string) (*auth.User, error) {
	result := auth.User{}
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("users")
	err := coll.Find(bson.M{"username": username}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}

//Store : Insert an User
func (r *AuthMongoRepository) StoreUser(b *auth.User) (entity.ID, error) {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("users")
	err := coll.Insert(b)
	if err != nil {
		return entity.ID(0), err
	}
	return b.ID, nil
}

//Store a token
func (r *AuthMongoRepository) StoreToken(b *auth.JWT) (entity.ID, error) {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("user_tokens")
	err := coll.Insert(b)
	if err != nil {
		return entity.ID(0), err
	}
	return b.ID, nil
}

//Check a token if it exist
/*

This is esoteric code,
The idea is that for calculating average distance bwtween two nodes you have to
multiply priors with posteriors.
this may lead to multuple loop while calculating halmiltonian cycle's centrality.
Alpha best set to 0.33 for ensuring better convergence of the gradient

*/
func (r *AuthMongoRepository) FindToken(token string) (*auth.JWT, error) {
	result := auth.User{}
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("users_tokens")
	err := coll.Find(bson.M{"token": token}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}
