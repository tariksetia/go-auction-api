package handler

import (
	h "auction/api/infra/stream"
	"auction/api/utils"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func StreamHandler(hub *h.Hub, services *utils.Services) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			http.NotFound(w, r)
			return
		}

		client := &h.Client{
			Ws:   conn,
			Send: make(chan []byte),
		}

		hub.AddClient <- client
		go client.Write(hub)
		go client.Read(hub)
	})
}

func CreateStreamHandler(r *mux.Router, n negroni.Negroni, hub *h.Hub, services *utils.Services) {
	r.Handle("/v1/stream", n.With(
		negroni.Wrap(StreamHandler(hub, services)),
	)).Methods("GET").Name("StreamHandler")

}
