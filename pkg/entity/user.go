package entity

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//User entity
type User struct {
	Id       ID     `json:"id" bson:"_id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//GenerateJWT
func (user User) GenerateJWT(key []byte) map[string]string {
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	userid := user.Id.String()
	claims["userID"] = userid
	claims["created"] = time.Now().Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(key)

	jwtMap := map[string]string{"token": tokenString}

	return jwtMap
}
