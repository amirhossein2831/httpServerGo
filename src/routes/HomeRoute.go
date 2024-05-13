package routes

import (
	"github.com/amirhossein2831/httpServerGo/src/http"
	"github.com/gorilla/mux"
)

func homeRoute(r *mux.Router) {
	r.HandleFunc("/home", http.HelloHandler).Methods("GET")
	r.HandleFunc("/form", http.FormHandler).Methods("GET")
}
