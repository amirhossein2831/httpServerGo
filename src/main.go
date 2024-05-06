package main

import (
	"github.com/amirhossein2831/httpServerGo/src/Http"
	"log"
	"net/http"
)

func main() {
	Http.Routing()

	println("server started at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
