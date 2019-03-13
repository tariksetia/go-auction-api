package middleware

import (
	"auction/api/infra/mongo"
	"auction/pkg/entity"
	"auction/pkg/user"
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	//Get and Parse Token data
	jwtToken := r.Context().Value("user").(*jwt.Token)
	created := jwtToken.Claims.(jwt.MapClaims)["created"].(float64)
	_id := jwtToken.Claims.(jwt.MapClaims)["userID"].(string)
	userID := entity.StringToID(_id)

	//Check if token is expired or not
	delta := time.Now().Unix() - int64(created)
	if delta > 3600 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Token Expired"))
		return
	}

	//conenct to db
	mpool, session := mongo.GetMongoPool(2)

	//create user repo and thus service
	userRepo := user.CreateMongoRepo(mpool, "auction")
	userService := user.NewService(userRepo)

	//get the user from userID
	usr, err := userService.Find(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Authentication Failed"))
		return
	}

	//Close the connection with db
	mpool.Close()
	session.Close()

	newRequest := r.WithContext(context.WithValue(r.Context(), "me", usr))
	next(w, newRequest)
}
