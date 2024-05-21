package request

import (
	"github.com/go-playground/validator/v10"
)

type CredRequest struct {
	Email    string `validate:"required,email" json:"email" `
	Password string `validate:"required,min=8" json:"password"`
}

func (c *CredRequest) Validate() (CredRequest, error) {
	validate := validator.New()
	err := validate.Struct(c)

	if err != nil {
		return CredRequest{}, validationError(err)
	}
	return *c, nil
}
