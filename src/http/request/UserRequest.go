package request

import (
	"github.com/amirhossein2831/httpServerGo/src/model"
	"github.com/go-playground/validator/v10"
)

type UserRequest struct {
	FirstName string `validate:"required" json:"first-name"`
	LastName  string `validate:"required" json:"last-name" `
	Email     string `validate:"required,email" json:"email" `
	Password  string `validate:"required,min=8" json:"password"`
}

func (ur *UserRequest) Validate() (model.User, error) {
	validate := validator.New()
	err := validate.Struct(ur)

	if err != nil {
		return model.User{}, validationError(err)
	}

	return model.User{
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		Email:     ur.Email,
		Password:  ur.Password,
	}, nil
}
