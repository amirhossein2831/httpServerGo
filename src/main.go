package main

import (
	"fmt"
	"github.com/amirhossein2831/httpServerGo/src/App"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/Logger"
	"github.com/amirhossein2831/httpServerGo/src/config"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"github.com/amirhossein2831/httpServerGo/src/routes"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
	// Init the app
	App.Configure()

	// defer the config
	defer App.DeferConfig()

	// Migrate Tables
	err := model.Migrate(DB.GetInstance().GetDb())
	if err != nil {
		Logger.GetInstance().GetLogger().Error("Database migration failed", zap.Error(err), zap.Time("timestamp", time.Now()))
		return
	}
	Logger.GetInstance().GetLogger().Info("Table Migrate successfully",
		zap.Time("timestamp", time.Now()),
	)

	// Start the App and server
	Logger.GetInstance().GetLogger().Info("App started successfully",
		zap.String("Port", config.GetInstance().Get("APP_PORT")),
		zap.Time("timestamp", time.Now()),
	)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", config.GetInstance().Get("APP_PORT")),
		routes.GetInstance().GetRouter()); err != nil {
		Logger.GetInstance().GetLogger().Error("Failed to start the server", zap.Error(err), zap.Time("timestamp", time.Now()))
	}
}
