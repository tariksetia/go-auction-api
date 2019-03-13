package bid

import "auction/pkg/entity"

//Repository repository interface
type Repository interface {
	Find(id entity.ID) (*Bid, error)
	Save(bid *Bid) (entity.ID, error)
	FindByKey(key string, val interface{}, page int, size int) ([]*Bid, error)
	Update(id entity.ID, key string, val interface{}) (*Bid, error)
}

//UseCase for offer
type UseCase interface {
	Find(id entity.ID) (*Bid, error)
	Save(bid *Bid) (entity.ID, error)
	FindByKey(key string, val interface{}, page int, size int) ([]*Bid, error)
	Update(id entity.ID, key string, val interface{}) (*Bid, error)
}
