package offers

import "time"
import "auction/pkg/entity"

type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Store an offer
func (s *Service) Store(offer *Offer) (entity.ID, error) {
	offer.ID = entity.NewID()
	offer.GoLive = time.Now()
	return s.repo.Store(offer)
}

//Find a offer by id
func (s *Service) Find(id entity.ID) (*Offer, error) {
	return s.repo.Find(id)
}

//query offer for pagination
func (s *Service) Query(page int, size int, sortkey string) ([]*Offer, error) {
	return s.repo.Query(page, size, sortkey)
}
