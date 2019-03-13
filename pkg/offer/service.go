package offer

import (
	"auction/pkg/entity"
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
func (s *Service) Save(user *Offer) (entity.ID, error) {
	user.ID = entity.NewID()
	return s.repo.Save(user)
}

//Find
func (s *Service) Find(id entity.ID) (*Offer, error) {
	return s.repo.Find(id)
}

//FindByKey
func (s *Service) FindByKey(key string, val interface{}, page int, size int) ([]*Offer, error) {
	return s.repo.FindByKey(key, val, page, size)
}

//FindByUsername
func (s *Service) Query(page int, size int, sortkey string) ([]*Offer, error) {
	return s.repo.Query(page, size, sortkey)
}

//Udapte
func (s *Service) Update(id entity.ID, key string, val interface{}) (*Offer, error) {
	return s.repo.Update(id, key, val)
}
