package routes

import (
	"github.com/amirhossein2831/httpServerGo/src/controller"
	"github.com/amirhossein2831/httpServerGo/src/util"
	"github.com/gorilla/mux"
	"net/http"
)

func Routing() *mux.Router {
	r := mux.NewRouter()

	// static file
	r.Handle("/", http.FileServer(http.Dir("static/html")))

	// home
	r.HandleFunc("/home", util.HelloHandler).Methods("GET")

	// user
	CrudRoute(r, "users", &controller.UserController{})

	return r
}
