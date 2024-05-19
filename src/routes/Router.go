package routes

import (
	"fmt"
	"github.com/amirhossein2831/httpServerGo/src/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func CrudRoute(r *mux.Router, path string, c controller.Crud, middlewares ...func(http.Handler) http.Handler) {
	subRouter := r.PathPrefix(fmt.Sprintf("/%v", path)).Subrouter()
	if middlewares != nil {
		for _, middleware := range middlewares {
			subRouter.Use(middleware)
		}
	}
	subRouter.HandleFunc("/", c.Index).Methods("GET")
	subRouter.HandleFunc("/", c.Create).Methods("POST")
	subRouter.HandleFunc("/{id}", c.Show).Methods("GET")
	subRouter.HandleFunc("/{id}", c.Update).Methods("PUT")
	subRouter.HandleFunc("/{id}", c.Delete).Methods("DELETE")
}

func Get(r *mux.Router, path string, handler func(http.ResponseWriter, *http.Request), middlewares ...func(http.Handler) http.Handler) {
	subRouter := r.PathPrefix("").Subrouter()
	if middlewares != nil {
		for _, middleware := range middlewares {
			subRouter.Use(middleware)
		}
	}
	subRouter.HandleFunc(path, handler).Methods("GET")
}

func Post(r *mux.Router, path string, handler func(http.ResponseWriter, *http.Request), middlewares ...func(http.Handler) http.Handler) {
	subRouter := r.PathPrefix("").Subrouter()
	if middlewares != nil {
		for _, middleware := range middlewares {
			subRouter.Use(middleware)
		}
	}
	subRouter.HandleFunc(path, handler).Methods("POST")
}

func Patch(r *mux.Router, path string, handler func(http.ResponseWriter, *http.Request), middlewares ...func(http.Handler) http.Handler) {
	subRouter := r.PathPrefix("").Subrouter()
	if middlewares != nil {
		for _, middleware := range middlewares {
			subRouter.Use(middleware)
		}
	}
	subRouter.HandleFunc(path, handler).Methods("PATCH")
}

func Put(r *mux.Router, path string, handler func(http.ResponseWriter, *http.Request), middlewares ...func(http.Handler) http.Handler) {
	subRouter := r.PathPrefix("").Subrouter()
	if middlewares != nil {
		for _, middleware := range middlewares {
			subRouter.Use(middleware)
		}
	}
	subRouter.HandleFunc(path, handler).Methods("PUT")
}

func Delete(r *mux.Router, path string, handler func(http.ResponseWriter, *http.Request), middlewares ...func(http.Handler) http.Handler) {
	subRouter := r.PathPrefix("").Subrouter()
	if middlewares != nil {
		for _, middleware := range middlewares {
			subRouter.Use(middleware)
		}
	}
	subRouter.HandleFunc(path, handler).Methods("DELETE")
}
