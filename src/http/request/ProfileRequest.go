package request

import (
	"github.com/amirhossein2831/httpServerGo/src/model"
	"github.com/go-playground/validator/v10"
	"time"
)

type ProfileRequest struct {
	Age       int       `validate:"required,min=0" json:"age"`
	Address   string    `validate:"required" json:"address"`
	BirthData time.Time `validate:"required" json:"birth-data"`
	UserID    uint      `validate:"" json:"user-id"`
}

func (p *ProfileRequest) Validate() (model.Profile, error) {
	validate := validator.New()
	err := validate.Struct(p)

	if err != nil {
		return model.Profile{}, validationError(err)
	}

	return model.Profile{
		Age: p.Age,
		Address:  p.Address,
		BirthData:     p.BirthData,
		UserID:  p.UserID,
	}, nil
}
