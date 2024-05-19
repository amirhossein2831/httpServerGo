package routes

import (
	"github.com/amirhossein2831/httpServerGo/src/Middleware"
	"github.com/amirhossein2831/httpServerGo/src/controller"
	"github.com/gorilla/mux"
	"net/http"
)

// see the api documentation in https://documenter.getpostman.com/view/29634924/2sA3JT1dNa

func Routing() *mux.Router {
	r := mux.NewRouter()
	subRouter := r.PathPrefix("/api/v1/").Subrouter()

	// static file
	r.Handle("/", http.FileServer(http.Dir("static/html")))
	r.Handle("/home", http.FileServer(http.Dir("static/html")))

	// single routes
	Post(subRouter, "/users/login/", controller.Login)

	// crud routes
	CrudRoute(subRouter, "users", &controller.UserController{}, Middleware.AuthMiddleware)
	CrudRoute(subRouter, "movies", &controller.MovieController{}, Middleware.AuthMiddleware)
	CrudRoute(subRouter, "books", &controller.BookController{}, Middleware.AuthMiddleware)

	return r
}
