package main

import (
	"github.com/amirhossein2831/httpServerGo/src/App"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"github.com/amirhossein2831/httpServerGo/src/routes"
	"log"
	"net/http"
)

func main() {
	// connect to DB
	db, err := DB.ConnectDB()
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
	App.App.Init(db, routes.Routing())

	// run server
	println("server started at port 8080")
	if err := http.ListenAndServe(":8080", App.App.GetRouter()); err != nil {
		log.Fatal(err)
	}
}
