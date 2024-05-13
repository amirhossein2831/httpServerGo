package Routes

import (
	"github.com/amirhossein2831/httpServerGo/src/Http"
	"github.com/gorilla/mux"
)

func homeRoute(r *mux.Router) {
	r.HandleFunc("/home", Http.HelloHandler).Methods("GET")
	r.HandleFunc("/form", Http.FormHandler).Methods("GET")
}
