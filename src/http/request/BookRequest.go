package request

import (
	"github.com/amirhossein2831/httpServerGo/src/model"
	"github.com/go-playground/validator/v10"
)

type BookRequest struct {
	Name        string `validate:"required,max=25" json:"name"`
	Author      string `validate:"required" json:"author"`
	Publication string `validate:"required" json:"publication"`
}

func (b *BookRequest) Validate() (model.Mod, error) {
	validate := validator.New()
	err := validate.Struct(b)

	if err != nil {
		return model.Book{}, validationError(err)
	}

	return model.Book{
		Name:        b.Name,
		Author:      b.Author,
		Publication: b.Publication,
	}, nil
}
