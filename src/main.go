package main

import (
	"fmt"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/config"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"github.com/amirhossein2831/httpServerGo/src/routes"
	"log"
	"net/http"
)

func main() {
	// migrate Tables
	err := model.Migrate(DB.GetInstance().GetDb())
	if err != nil {
		return
	}
	println("Table Migrate successfully")

	router := routes.Routing()

	// run server
	println(fmt.Sprintf("app started at port %v", config.GetInstance().Get("APP_PORT")))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", config.GetInstance().Get("APP_PORT")), router); err != nil {
		log.Fatal(err)
	}
}
