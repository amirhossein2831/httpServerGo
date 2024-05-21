package service

import (
	"github.com/amirhossein2831/httpServerGo/src/http/repositories"
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/model"
)

type MovieService struct {
	repository repositories.Repository
}

func NewMovieService() *MovieService {
	return &MovieService{
		repository: repositories.NewMovieRepository(),
	}
}

func (s *MovieService) Index() ([]model.Mod, error) {
	movies, err := s.repository.All()
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (s *MovieService) Show(id string) (model.Mod, error) {
	movie, err := s.repository.Get(id)
	if err != nil {
		return model.Movie{}, err
	}
	return movie, nil
}

func (s *MovieService) Create(request request.Request) (model.Mod, error) {
	movie, err := request.Validate()
	if err != nil {
		return model.Movie{}, err
	}
	movie, err = s.repository.Create(movie)
	if err != nil {
		return model.Movie{}, err

	}
	return movie, err
}

func (s *MovieService) Update(request request.Request, id string) (model.Mod, error) {
	movie, err := request.Validate()
	if err != nil {
		return model.Movie{}, err
	}

	movie, err = s.repository.Update(movie, id)
	if err != nil {
		return model.Movie{}, err
	}
	return movie, nil
}

func (s *MovieService) Delete(id string) error {
	err := s.repository.SoftDelete(id)
	if err != nil {
		return err
	}
	return nil

}
