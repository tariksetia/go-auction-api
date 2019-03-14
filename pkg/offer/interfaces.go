package offer

import (
	e "auction/pkg/entity"
)

//Repository repository interface
type Repository interface {
	Find(id e.ID) (*e.Offer, error)
	Save(offer *e.Offer) (e.ID, error)
	FindByKey(key string, val interface{}, page int, size int) ([]*e.Offer, error)
	Query(page int, size int, sortkey string) ([]*e.Offer, error)
	Update(id e.ID, key string, val interface{}) (*e.Offer, error)
}

//UseCase for offer
type UseCase interface {
	Find(id e.ID) (*e.Offer, error)
	Save(user *e.Offer) (e.ID, error)
	FindByKey(key string, val interface{}, page int, size int) ([]*e.Offer, error)
	Query(page int, size int, sortkey string) ([]*e.Offer, error)
	Update(id e.ID, key string, val interface{}) (*e.Offer, error)
}
