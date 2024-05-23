package main

import (
	"fmt"
	"github.com/amirhossein2831/httpServerGo/src/App"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/config"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"github.com/amirhossein2831/httpServerGo/src/routes"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

func main() {
	// init the app
	App.Configure()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// migrate Tables
	err := model.Migrate(DB.GetInstance().GetDb())
	if err != nil {
		logger.Error("Database migration failed", zap.Error(err))
		return
	}
	logger.Info("Table Migrate successfully",
		zap.Time("timestamp", time.Now()),
	)

	println("Table Migrate successfully")

	// run server
	println(fmt.Sprintf("app started at port %v", config.GetInstance().Get("APP_PORT")))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", config.GetInstance().Get("APP_PORT")),
		routes.GetInstance().GetRouter()); err != nil {
		log.Fatal(err)
	}
}
