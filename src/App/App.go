package App

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var App = &APP{}

type APP struct {
	db *gorm.DB
	r  *mux.Router
}

func (a *APP) Init(db *gorm.DB, r *mux.Router) {
	a.db = db
	a.r = r
}

func (a *APP) GetDB() *gorm.DB {
	return a.db
}

func (a *APP) GetRouter() *mux.Router {
	return a.r
}
