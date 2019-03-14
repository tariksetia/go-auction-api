package user

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
func (s *Service) Save(user *e.User) (e.ID, error) {
	user.Id = e.NewID()
	return s.repo.Save(user)
}

//Find
func (s *Service) Find(id e.ID) (*e.User, error) {
	return s.repo.Find(id)
}

//FindByKey
func (s *Service) FindByKey(key string, val interface{}) ([]*e.User, error) {
	return s.repo.FindByKey(key, val)
}

//FindByUsername
func (s *Service) FindByUsername(username string) ([]*e.User, error) {
	return s.repo.FindByKey("username", username)
}
