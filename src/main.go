package main

import (
	"github.com/amirhossein2831/httpServerGo/src/config"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"github.com/amirhossein2831/httpServerGo/src/routes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	// load env variable
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// connect to DB
	config.App.DB, err = config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// instantiate router
	config.App.R = routes.Routing()

	// seed data
	model.SeedUser()

	// run server
	println("Connected to DB successfully")
	println("server started at port 8080")
	if err := http.ListenAndServe(":8080", config.App.R); err != nil {
		log.Fatal(err)
	}
}
