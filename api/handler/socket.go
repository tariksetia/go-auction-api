package handler

import (
	h "auction/api/infra/hub"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func StreamHandler(hub *h.Hub) http.Handler {
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
		go client.Write()
		go client.Read()
	})
}

func CreateStreamHandler(r *mux.Router, n negroni.Negroni, hub *h.Hub) {
	r.Handle("/v1/stream", n.With(
		negroni.Wrap(StreamHandler(hub)),
	)).Methods("GET").Name("StreamHandler")

}
