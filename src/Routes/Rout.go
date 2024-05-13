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
	r.HandleFunc("/users/", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users/", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	return r
}
