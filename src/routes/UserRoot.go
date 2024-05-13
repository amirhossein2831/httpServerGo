package routes

import (
	"github.com/amirhossein2831/httpServerGo/src/http"
	"github.com/gorilla/mux"
)

func userRoute(r *mux.Router) {
	r.HandleFunc("/users/", http.GetUsers).Methods("GET")
	r.HandleFunc("/users/", http.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", http.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", http.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", http.DeleteUser).Methods("DELETE")
}
