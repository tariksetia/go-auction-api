package offer

import (
	"auction/pkg/entity"
)

//Repository repository interface
type Repository interface {
	Find(id entity.ID) (*Offer, error)
	Save(offer *Offer) (entity.ID, error)
	FindByKey(key string, val interface{}, page int, size int) ([]*Offer, error)
	Query(page int, size int, sortkey string) ([]*Offer, error)
	Update(id entity.ID, key string, val interface{}) (*Offer, error)
}

//UseCase for offer
type UseCase interface {
	Find(id entity.ID) (*Offer, error)
	Save(user *Offer) (entity.ID, error)
	FindByKey(key string, val interface{}, page int, size int) ([]*Offer, error)
	Query(page int, size int, sortkey string) ([]*Offer, error)
	Update(id entity.ID, key string, val interface{}) (*Offer, error)
}
