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
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	println("Connected to DB successfully")

	// migrate Tables
	err = model.Migrate(db)
	if err != nil {
		return
	}
	println("Table Migrate successfully")

	// Init App
	config.App.Init(db, routes.Routing())

	// run server
	println("server started at port 8080")
	if err := http.ListenAndServe(":8080", config.App.GetRouter()); err != nil {
		log.Fatal(err)
	}
}
