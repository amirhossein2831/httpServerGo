package request

import "github.com/amirhossein2831/httpServerGo/src/model"

type UserRequest struct {
	FirstName string        `json:"first-name"`
	LastName  string        `json:"last-name"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	Profile   model.Profile `json:"profile"`
}

func (ur *UserRequest) ToUser() model.User {
	return model.User{
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		Email:     ur.Email,
		Password:  ur.Password,
		Profile:   ur.Profile,
	}
}
