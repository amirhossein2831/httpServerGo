package Http

import "net/http"

func Routing() {
	http.Handle("/", http.FileServer(http.Dir("static/html")))
	http.HandleFunc("/home", HelloHandler)
	http.HandleFunc("/form", FormHandler)
}
