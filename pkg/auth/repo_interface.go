package auth

import "auction/pkg/entity"

type Repository interface {
	FindByUserid(id entity.ID) (*User, error)
	FindByUsername(username string) (*User, error)
	FindByToken(token string) (*JWT, error)
	StoreUser(user *User) (entity.ID, error)
	FindToken(token string) (entity.ID, error)
}
