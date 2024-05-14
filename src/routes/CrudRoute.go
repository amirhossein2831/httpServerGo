package routes

import (
	"fmt"
	"github.com/amirhossein2831/httpServerGo/src/controller"
	"github.com/gorilla/mux"
)

func CrudRoute(r *mux.Router, path string, c controller.Crud) {
	r.HandleFunc(fmt.Sprintf("/%v/", path), c.Index).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%v/", path), c.Create).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%v/{id}", path), c.Show).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%v/{id}", path), c.Update).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%v/{id}", path), c.Delete).Methods("DELETE")
}
