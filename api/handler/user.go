package handler

import (
	"auction/api/config"
	e "auction/pkg/entity"
	"auction/pkg/user"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func signup(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error sigining up user"
		var usr *e.User

		err := json.NewDecoder(r.Body).Decode(&usr)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := service.FindByUsername(usr.Username)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if data != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("User Already Exist"))
			return
		}
		usr.Password = user.SaltPassowrd(usr.Password)
		usr.Id, err = service.Save(usr)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusCreated)

	})
}

func login(service user.UseCase) http.Handler {
	cfg := config.GetAppConfig()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading user"
		var usr *e.User
		err := json.NewDecoder(r.Body).Decode(&usr)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.FindByUsername(usr.Username)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("User Doesn't Exist"))
			return
		}

		userDatum := data[0]

		if !user.ComparePasswords(userDatum.Password, usr.Password) {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Password Doesnot Match"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != e.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		jwtmap := userDatum.GenerateJWT([]byte(cfg.GetAppSecret()))
		if err := json.NewEncoder(w).Encode(jwtmap); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusAccepted)
	})

}

//CreateUserHandlers Maps routes to http handlers
func CreateUserHandlers(r *mux.Router, n negroni.Negroni, service user.UseCase) {
	r.Handle("/v1/login", n.With(
		negroni.Wrap(login(service)),
	)).Methods("POST", "OPTIONS").Name("login")

	r.Handle("/v1/signup", n.With(
		negroni.Wrap(signup(service)),
	)).Methods("POST", "OPTIONS").Name("signup")
}
