package routes

import (
	"github.com/amirhossein2831/httpServerGo/src/Logger"
	"github.com/amirhossein2831/httpServerGo/src/http/Middleware"
	"github.com/amirhossein2831/httpServerGo/src/http/controller"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"sync"
	"time"
)

// see the api documentation in https://documenter.getpostman.com/view/29634924/2sA3JT1dNa
var (
	routerInstance Route
	once           sync.Once
)

type Route interface {
	GetRouter() *mux.Router
	SetRouter(router *mux.Router)
	Routing()
}

type Router struct {
	r *mux.Router
}

func (r *Router) GetRouter() *mux.Router {
	return r.r
}

func (r *Router) SetRouter(router *mux.Router) {
	r.r = router
}

func GetInstance() Route {
	once.Do(func() {
		routerInstance = &Router{r: mux.NewRouter()}
		routerInstance.Routing()
		Logger.GetInstance().GetLogger().Info("router instantiate successfully",
			zap.Time("timestamp", time.Now()),
		)
	})
	return routerInstance
}

func (r *Router) Routing() {
	// log all incoming request
	r.r.Use(Middleware.Log)

	// create api/v1 route group
	subRouter := r.r.PathPrefix("/api/v1/").Subrouter()

	// static file
	r.r.Handle("/", http.FileServer(http.Dir("static/html")))
	r.r.Handle("/home", http.FileServer(http.Dir("static/html")))

	// single routes
	Post(subRouter, "/users/login/", controller.Login)

	// crud routes
	CrudRoute(subRouter, "users", controller.NewUserController(), Middleware.CustomPermission(map[string][]string{
		"GET":    {"read-user"},
		"PUT":    {"update-user"},
		"DELETE": {"delete-user"},
	}))
	CrudRoute(subRouter, "movies", controller.NewMovieController(), Middleware.Authenticate, Middleware.Role([]string{"movie"}))
	CrudRoute(subRouter, "books", controller.NewBookController(), Middleware.Authenticate, Middleware.Role([]string{"book"}))
}
