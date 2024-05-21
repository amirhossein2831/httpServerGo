package service

import (
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/model"
)

type Service interface {
	Index() ([]model.Mod, error)
	Show(string) (model.Mod, error)
	Create(request.UserRequest) (model.Mod, error)
	Update(request.UserRequest, string) (model.Mod, error)
	Delete(string) error
}
