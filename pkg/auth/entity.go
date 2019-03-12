package auth

import (
	"auction/pkg/entity"
)

//User entity
type User struct {
	ID       entity.ID
	Username string
	password string
}

//JWT enttiy table
type JWT struct {
	ID  entity.ID
	jwt string
}
