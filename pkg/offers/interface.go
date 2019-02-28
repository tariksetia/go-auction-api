package offers

import "auction/pkg/entity"

//Repository repository interface
type Repository interface {
	Find(id entity.ID) (*Offer, error)
	Store(offer *Offer) (entity.ID, error)
	Query(page int, size int, sortkey string) ([]*Offer, error)
}
