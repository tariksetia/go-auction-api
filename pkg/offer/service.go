package offer

import (
	e "auction/pkg/entity"
)

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Save
func (s *Service) Save(offer *e.Offer) (e.ID, error) {
	offer.Id = e.NewID()
	return s.repo.Save(offer)
}

//Find
func (s *Service) Find(id e.ID) (*e.Offer, error) {
	return s.repo.Find(id)
}

//FindByKey
func (s *Service) FindByKey(key string, val interface{}, page int, size int) ([]*e.Offer, error) {
	return s.repo.FindByKey(key, val, page, size)
}

//FindByUsername
func (s *Service) Query(page int, size int, sortkey string) ([]*e.Offer, error) {
	return s.repo.Query(page, size, sortkey)
}

//Udapte
func (s *Service) Update(id e.ID, key string, val interface{}) (*e.Offer, error) {
	return s.repo.Update(id, key, val)
}
