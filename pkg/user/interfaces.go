package user

import "auction/pkg/entity"

type Repository interface {
	Find(id entity.ID) (*User, error)
	Save(user *User) (entity.ID, error)
	FindByKey(key string, val interface{}) ([]*User, error)
}

type UseCase interface {
	Find(id entity.ID) (*User, error)
	Save(user *User) (entity.ID, error)
	FindByKey(key string, val interface{}) ([]*User, error)
	FindByUsername(username string) ([]*User, error)
}
