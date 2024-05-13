package main

import (
	"github.com/amirhossein2831/httpServerGo/src/Model"
	"github.com/amirhossein2831/httpServerGo/src/Routes"
	"github.com/amirhossein2831/httpServerGo/src/config"
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
	DB, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// seed data
	Model.SeedUser()

	// instantiate router
	router := Routes.Routing()

	// run server
	println("Connected to DB successfully")
	println("server started at port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
	print(DB.Error)
}
