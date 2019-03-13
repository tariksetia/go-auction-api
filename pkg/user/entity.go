package user

import "auction/pkg/entity"

//User entity
type User struct {
	ID       entity.ID `json:"id" bson:"_id,omitempty"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
