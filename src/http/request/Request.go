package request

import (
	"errors"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"github.com/go-playground/validator/v10"
)

type Request interface {
	Validate() (model.Mod, error)
}

func validationError(err error) error {
	var validationErrors validator.ValidationErrors
	errors.As(err, &validationErrors)
	valError := ""
	for _, err := range validationErrors {
		valError += err.Field() + ": this field has error: " + err.Tag() + " " + err.Param() + " | "
	}
	return errors.New(valError)
}
