package main

import (
	"github.com/amirhossein2831/httpServerGo/src/Http"
	"github.com/amirhossein2831/httpServerGo/src/Model"
	"log"
	"net/http"
)

func main() {
	// seed data
	Model.SeedUser()

	// instantiate router
	router := Http.Routing()

	// run server
	println("server started at port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
