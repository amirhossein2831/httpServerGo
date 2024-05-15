package routes

import (
	"github.com/amirhossein2831/httpServerGo/src/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func Routing() *mux.Router {
	r := mux.NewRouter()

	// static file
	r.Handle("/", http.FileServer(http.Dir("static/html")))

	// home
	r.HandleFunc("/home", controller.Index).Methods("GET")

	// user
	CrudRoute(r, "users", &controller.UserController{})
	CrudRoute(r, "movies", &controller.MovieController{})
	CrudRoute(r, "books", &controller.BookController{})

	return r
}
