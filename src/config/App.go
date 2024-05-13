package config

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type APP struct {
	DB *gorm.DB
	R  *mux.Router
}

var App = &APP{}
