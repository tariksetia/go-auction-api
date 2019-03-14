package user

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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
