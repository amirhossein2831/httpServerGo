package routes

import (
	"github.com/amirhossein2831/httpServerGo/src/controller"
	"github.com/gorilla/mux"
	"net/http"
)

// see the api documentation in https://documenter.getpostman.com/view/29634924/2sA3JRYeTg

func Routing() *mux.Router {
	r := mux.NewRouter()
	subRouter := r.PathPrefix("/api/v1/").Subrouter()

	// static file
	r.Handle("/", http.FileServer(http.Dir("static/html")))
	r.Handle("/home", http.FileServer(http.Dir("static/html")))

	// user
	CrudRoute(subRouter, "users", &controller.UserController{})
	CrudRoute(subRouter, "movies", &controller.MovieController{})
	CrudRoute(subRouter, "books", &controller.BookController{})

	return r
}
