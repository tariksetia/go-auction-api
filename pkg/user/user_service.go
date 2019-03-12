package user

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
func (s *Service) Save(user *User) (entity.ID, error) {
	user.ID = entity.NewID()
	return s.repo.Save(user)
}

//Find
func (s *Service) Find(id entity.ID) (*User, error) {
	return s.repo.Find(id)
}

//FindByKey
func (s *Service) FindByKey(key string, val interface{}) ([]*User, error) {
	return s.repo.FindByKey(key, val)
}

//FindByUsername
func (s *Service) FindByUsername(username string) ([]*User, error) {
	return s.repo.FindByKey("username", username)
}
