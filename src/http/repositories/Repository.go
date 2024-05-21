package repositories

import "github.com/amirhossein2831/httpServerGo/src/model"

type Repository interface {
	All() ([]model.Mod, error)
	Get(string) (model.Mod, error)
	GetByColumn(string, string) (model.Mod, error)
	Create(model.Mod) (model.Mod, error)
	Update(model.Mod, string) (model.Mod, error)
	SoftDelete(string) error
	HardDelete(string) error
}
