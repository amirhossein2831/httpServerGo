package main

import (
	"fmt"
	"github.com/amirhossein2831/httpServerGo/src/App"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/config"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"github.com/amirhossein2831/httpServerGo/src/routes"
	"log"
	"net/http"
)

func main() {
	// init the app
	App.Configure()

	// migrate Tables
	err := model.Migrate(DB.GetInstance().GetDb())
	if err != nil {
		return
	}
	println("Table Migrate successfully")

	println(fmt.Sprintf("app started at port %v", config.GetInstance().Get("APP_PORT")))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", config.GetInstance().Get("APP_PORT")),
		routes.GetInstance().GetRouter()); err != nil {
		log.Fatal(err)
	}
}

// create an app with Db,config,router
// it should not be single ton because we can have several app
