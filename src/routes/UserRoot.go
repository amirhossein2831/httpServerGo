package routes

import (
	"github.com/amirhossein2831/httpServerGo/src/controller"
	"github.com/gorilla/mux"
)

func userRoute(r *mux.Router) {
	r.HandleFunc("/users/", controller.GetUsers).Methods("GET")
	r.HandleFunc("/users/", controller.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controller.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", controller.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")
}
