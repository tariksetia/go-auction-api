package user

import e "auction/pkg/entity"

type Repository interface {
	Find(id e.ID) (*e.User, error)
	Save(user *e.User) (e.ID, error)
	FindByKey(key string, val interface{}) ([]*e.User, error)
}

type UseCase interface {
	Find(id e.ID) (*e.User, error)
	Save(user *e.User) (e.ID, error)
	FindByKey(key string, val interface{}) ([]*e.User, error)
	FindByUsername(username string) ([]*e.User, error)
}
