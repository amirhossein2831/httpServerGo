package service

import (
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"github.com/amirhossein2831/httpServerGo/src/repositories"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Get() ([]model.User, error) {
	users, err := repositories.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetEntity(id string) (model.User, error) {
	users, err := repositories.GetUserById(id)
	if err != nil {
		return model.User{}, err
	}
	return users, nil
}

func (s *UserService) Post(request request.UserRequest) (model.User, error) {
	user, err := request.Validate()
	if err != nil {
		return model.User{}, err
	}
	user, err = repositories.CreateUser(user)
	if err != nil {
		return model.User{}, err

	}
	return user, err
}

func (s *UserService) Put(request request.UserRequest, id string) (model.User, error) {
	user, err := request.Validate()
	if err != nil {
		return model.User{}, err
	}

	user, err = repositories.UpdateUser(user, id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *UserService) Delete(id string) error {
	err := repositories.SoftDeleteUser(id)
	if err != nil {
		return err
	}
	return nil

}
