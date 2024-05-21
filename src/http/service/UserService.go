package service

import (
	"github.com/amirhossein2831/httpServerGo/src/http/repositories"
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/model"
)

type UserService struct {
	repository repositories.Repository
}

func NewUserService() *UserService {
	return &UserService{
		repository: repositories.NewUserRepository(),
	}
}

func (s *UserService) Index() ([]model.Mod, error) {
	users, err := s.repository.All()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) Show(id string) (model.Mod, error) {
	user, err := s.repository.Get(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *UserService) Create(request request.Request) (model.Mod, error) {
	user, err := request.Validate()
	if err != nil {
		return model.User{}, err
	}
	user, err = s.repository.Create(user)
	if err != nil {
		return model.User{}, err

	}
	return user, err
}

func (s *UserService) Update(request request.Request, id string) (model.Mod, error) {
	user, err := request.Validate()
	if err != nil {
		return model.User{}, err
	}

	user, err = s.repository.Update(user, id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *UserService) Delete(id string) error {
	err := s.repository.SoftDelete(id)
	if err != nil {
		return err
	}
	return nil

}
