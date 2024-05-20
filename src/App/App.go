package App

import (
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/config"
	"github.com/amirhossein2831/httpServerGo/src/routes"
)

var application APP

type APP interface {
	Db() DB.Database
	SetDb(DB.Database)
	Router() routes.Route
	SetRouter(routes.Route)
	Config() config.Configurator
	SetConfig(config.Configurator)
}

type App struct {
	db     DB.Database
	router routes.Route
	config config.Configurator
}

func Configure() {
	application = &App{
		db:     DB.GetInstance(),
		router: routes.GetInstance(),
		config: config.GetInstance(),
	}
}

func GetApp() APP {
	return application
}

func (app *App) Db() DB.Database {
	return app.db
}

func (app *App) SetDb(db DB.Database) {
	app.db = db
}

func (app *App) Router() routes.Route {
	return app.router
}

func (app *App) SetRouter(router routes.Route) {
	app.router = router
}

func (app *App) Config() config.Configurator {
	return app.config
}

func (app *App) SetConfig(config config.Configurator) {
	app.config = config
}
