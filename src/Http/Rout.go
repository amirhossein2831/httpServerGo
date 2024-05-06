package Http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Routing() *mux.Router {
	r := mux.NewRouter()

	// static file
	http.Handle("/", http.FileServer(http.Dir("static/html")))

	// home
	r.HandleFunc("/home", HelloHandler)
	r.HandleFunc("/form", FormHandler)

	// movie
	r.HandleFunc("/movies", MoviesHandler)
	r.HandleFunc("/movies/{id}", MovieHandler)

	return r
}

