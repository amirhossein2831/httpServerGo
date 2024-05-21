package service

import (
	"github.com/amirhossein2831/httpServerGo/src/http/repositories"
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/model"
)

type BookService struct {
	repository repositories.Repository
}

func NewBookService() *BookService {
	return &BookService{
		repository: repositories.NewBookRepository(),
	}
}

func (s *BookService) Index() ([]model.Mod, error) {
	books, err := s.repository.All()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BookService) Show(id string) (model.Mod, error) {
	book, err := s.repository.Get(id)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (s *BookService) Create(request request.Request) (model.Mod, error) {
	book, err := request.Validate()
	if err != nil {
		return model.Book{}, err
	}
	book, err = s.repository.Create(book)
	if err != nil {
		return model.Book{}, err

	}
	return book, err
}

func (s *BookService) Update(request request.Request, id string) (model.Mod, error) {
	book, err := request.Validate()
	if err != nil {
		return model.Book{}, err
	}

	book, err = s.repository.Update(book, id)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (s *BookService) Delete(id string) error {
	err := s.repository.SoftDelete(id)
	if err != nil {
		return err
	}
	return nil

}
