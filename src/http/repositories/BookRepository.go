package repositories

import (
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"strconv"
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

func CreateBook(book model.Book) (model.Book, error) {
	res := DB.GetInstance().GetDb().Create(&book)
	if res.Error != nil {
		return model.Book{}, res.Error
	}

	return book, nil
}

func UpdateBook(book model.Book, id string) (model.Book, error) {
	err := HardDeleteBook(id)
	if err != nil {
		return model.Book{}, err
	}

	Id, _ := strconv.Atoi(id)
	book.ID = uint(Id)
	res := DB.GetInstance().GetDb().Create(&book)
	if res.Error != nil {
		return model.Book{}, res.Error
	}

	return book, nil
}

func SoftDeleteBook(id string) error {
	if err := DB.GetInstance().GetDb().Delete(&model.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}

func HardDeleteBook(id string) error {
	if err := DB.GetInstance().GetDb().Unscoped().Delete(&model.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}
