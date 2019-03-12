package user

import (
	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//As a standard practice we store hashed version of Password
func SaltPassowrd(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Error(err)
	}
	return string(hash)
}

//Compare the the password with hash passed
func ComparePasswords(hashVal string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(password))
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

//GenerateJWT
func GenerateJWT(user *User, key []byte) map[string]string {
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["user"] = user
	claims["userID"] = user.ID
	claims["created"] = string(time.Now().Unix())

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(key)

	jwtMap := map[string]string{"token": tokenString}

	return jwtMap
}
