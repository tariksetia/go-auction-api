package bid

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
func (s *Service) Save(user *Bid) (entity.ID, error) {
	user.ID = entity.NewID()
	return s.repo.Save(user)
}

//Find
func (s *Service) Find(id entity.ID) (*Bid, error) {
	return s.repo.Find(id)
}

//FindByKey
func (s *Service) FindByKey(key string, val interface{}, page int, size int) ([]*Bid, error) {
	return s.repo.FindByKey(key, val, page, size)
}

//Udapte
func (s *Service) Update(id entity.ID, key string, val interface{}) (*Bid, error) {
	return s.repo.Update(id, key, val)
}
