package repositories

import (
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"net/http"
)

func GetBooks() ([]model.Book, error) {
	var books []model.Book
	res := DB.GetInstance().GetDb().Find(&books)
	if res.Error != nil {
		return nil, res.Error
	}
	return books, nil
}

func GetBook(id string) (model.Book, error) {
	var book model.Book

	res := DB.GetInstance().GetDb().First(&book, id)
	if res.Error != nil {
		return book, res.Error
	}
	return book, nil
}

func CreateBook(r *http.Request) (model.Book, error) {
	var bookRequest request.BookRequest
	err := json.NewDecoder(r.Body).Decode(&bookRequest)
	if err != nil {
		return model.Book{}, err
	}
	book := bookRequest.ToBook()

	res := DB.GetInstance().GetDb().Create(&book)
	if res.Error != nil {
		return model.Book{}, res.Error
	}

	return book, nil
}

func UpdateBook(r *http.Request, id string) (model.Book, error) {
	var bookRequest request.BookRequest
	err := DeleteBook(id)
	if err != nil {
		return model.Book{}, err
	}
	err = json.NewDecoder(r.Body).Decode(&bookRequest)
	if err != nil {
		return model.Book{}, err
	}
	book := bookRequest.ToBook()

	res := DB.GetInstance().GetDb().Create(&book)
	if res.Error != nil {
		return model.Book{}, res.Error
	}

	return book, nil
}

func DeleteBook(id string) error {
	if err := DB.GetInstance().GetDb().Delete(&model.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}
