package request

import (
	"github.com/amirhossein2831/httpServerGo/src/model"
	"github.com/go-playground/validator/v10"
)

type MovieRequest struct {
	Name        string `validate:"required" json:"name"`
	Director    string `validate:"required" json:"director"`
	Publication string `validate:"required" json:"publication"`
	WatchTime   int    `validate:"required,min=0" json:"watch_time"`
}

func (m *MovieRequest) Validate() (model.Mod, error) {
	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		return model.Movie{}, validationError(err)
	}

	return model.Movie{
		Name:        m.Name,
		Director:    m.Director,
		Publication: m.Publication,
		WatchTime:   m.WatchTime,
	}, nil
}
