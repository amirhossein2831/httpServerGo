package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Routing() *mux.Router {
	r := mux.NewRouter()

	// static file
	r.Handle("/", http.FileServer(http.Dir("static/html")))
	// home
	homeRoute(r)
	// user
	userRoute(r)

	return r
}
