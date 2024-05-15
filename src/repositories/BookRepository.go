package repositories

import (
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/config"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"net/http"
)

func GetBooks() ([]model.Book, error) {
	var books []model.Book
	res := config.App.GetDB().Find(&books)
	if res.Error != nil {
		return nil, res.Error
	}
	return books, nil
}

func GetBook(id string) (model.Book, error) {
	var book model.Book

	res := config.App.GetDB().First(&book, id)
	if res.Error != nil {
		return book, res.Error
	}
	return book, nil
}

func CreateBook(r *http.Request) (model.Book, error) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		return book, err
	}
	res := config.App.GetDB().Create(&book)
	if res.Error != nil {
		return book, res.Error
	}

	return book, nil
}

func UpdateBook(r *http.Request, id string) (model.Book, error) {
	var book model.Book
	err := config.App.GetDB().First(&book, id).Error
	if err != nil {
		return book, err
	}
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		return book, err
	}
	config.App.GetDB().Save(&book)

	return book, nil
}

func DeleteBook(id string) error {
	if err := config.App.GetDB().Delete(&model.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}
