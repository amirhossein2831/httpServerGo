package repositories

import "github.com/amirhossein2831/httpServerGo/src/model"

type Repository interface {
	All() ([]model.User, error)
	Get(string) (model.User, error)
	GetByColumn(string) (model.User, error)
	Create(model.User) (model.User, error)
	Update(model.User, string) (model.User, error)
	SoftDelete(string) error
	HardDelete(string) error
}
