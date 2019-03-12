package auth

import "auction/pkg/entity"

type Service struct {
	repo Repository
}

//NewService create new service
func CreateAuthService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Create a user
func (s *Service) SignUp(user *User) (entity.ID, error) {

	user.ID = entity.NewID()
	return s.repo.StoreUser(user)
}

//Get a user info by user name
func (s *Service) FindByUserName(username string) (*User, error) {
	return s.repo.FindByUsername(username)
}
