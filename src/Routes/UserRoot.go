package Routes

import (
	"github.com/amirhossein2831/httpServerGo/src/Http"
	"github.com/gorilla/mux"
)

func userRoute(r *mux.Router) {
	r.HandleFunc("/users/", Http.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", Http.GetUser).Methods("GET")
	r.HandleFunc("/users/", Http.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", Http.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", Http.DeleteUser).Methods("DELETE")
}
