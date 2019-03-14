package bid

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
func (s *Service) Save(bid *e.Bid) (e.ID, error) {
	bid.Id = e.NewID()
	return s.repo.Save(bid)
}

//Find
func (s *Service) Find(id e.ID) (*e.Bid, error) {
	return s.repo.Find(id)
}

//FindByKey
func (s *Service) FindByKey(key string, val interface{}, page int, size int) ([]*e.Bid, error) {
	return s.repo.FindByKey(key, val, page, size)
}

//Udapte
func (s *Service) Update(id e.ID, key string, val interface{}) (*e.Bid, error) {
	return s.repo.Update(id, key, val)
}
