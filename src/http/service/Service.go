package service

import (
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/model"
)

type Service interface {
	Index() ([]model.User, error)
	Show(string) (model.User, error)
	Create(request.UserRequest) (model.User, error)
	Update(request.UserRequest, string) (model.User, error)
	Delete(string) error
}
