package Http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Routing() *mux.Router {
	r := mux.NewRouter()

	// static file
	r.Handle("/", http.FileServer(http.Dir("static/html")))

	// home
	r.HandleFunc("/home", helloHandler).Methods("GET")
	r.HandleFunc("/form", formHandler).Methods("GET")

	// movie
	r.HandleFunc("/movies", allUser).Methods("GET")
	r.HandleFunc("/movies/{id}", getUser).Methods("GET")
	r.HandleFunc("/movies/", createUser).Methods("POST")
	r.HandleFunc("/movies/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/movies/{id}", partialUpdateUser).Methods("PATCH")
	r.HandleFunc("/movies/{id}", deleteUser).Methods("DELETE")

	return r
}
