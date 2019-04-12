package utils

import (
	"auction/api/config"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

func ParseAuthToken(tokenstring string) (jwt.Claims, error) {
	cfg := config.GetAppConfig()

	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.GetAppSecret()), nil
	})

	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}

func ValidateAuthToken(tokenstring string) (int,error) {
	claims, err := ParseAuthToken(tokenstring)
	if err != nil {
		return 0,errors.New("Error While parsing security token")
	}
	created := claims.(jwt.MapClaims)["created"].(float64)
	//Check if token is expired or not
	delta := time.Now().Unix() - int64(created)
	if delta > 3600 {
		return 0,errors.New("Token Expired")
	}
	return int(delta),nil
}
