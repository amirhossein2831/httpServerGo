package service

import (
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/model"
)

type Service interface {
	Index() ([]model.Mod, error)
	Show(string) (model.Mod, error)
	Create(request.Request) (model.Mod, error)
	Update(request.Request, string) (model.Mod, error)
	Delete(string) error
}
