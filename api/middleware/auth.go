package middleware

import (
	"auction/api/config"
	"auction/pkg/entity"
	"auction/pkg/user"
	"context"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

//LoginMiddleware Create Negroni based login middle, requre resference to userService as pointer
func LoginMiddleware(userService *user.Service) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

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

		//get the user from userID
		usr, err := userService.Find(userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Authentication Failed"))
			return
		}

		newRequest := r.WithContext(context.WithValue(r.Context(), "me", usr))
		next(w, newRequest)
	})

}

func JwtMiddleware(cfg *config.AppConfig) negroni.HandlerFunc {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.GetAppSecret()), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return jwtMiddleware.HandlerWithNext
}
