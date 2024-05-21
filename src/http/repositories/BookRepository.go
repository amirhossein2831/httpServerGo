package repositories

import (
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"strconv"
)

type BookRepository struct {
}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (ur *BookRepository) All() ([]model.Mod, error) {
	var books []model.Book

	err := DB.GetInstance().GetDb().Find(&books).Error
	if err != nil {
		return nil, err
	}

	return model.BookToMod(books), nil
}

func (ur *BookRepository) Get(id string) (model.Mod, error) {
	var book model.Book

	err := DB.GetInstance().GetDb().First(&book, id).Error
	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}
func (ur *BookRepository) GetByColumn(column, value string) (model.Mod, error) {
	return nil, nil
}

func (ur *BookRepository) Create(data model.Mod) (model.Mod, error) {
	book := data.(model.Book)

	err := DB.GetInstance().GetDb().Create(&book).Error
	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (ur *BookRepository) Update(data model.Mod, id string) (model.Mod, error) {
	book := data.(model.Book)
	err := ur.HardDelete(id)
	if err != nil {
		return model.Book{}, err
	}

	Id, _ := strconv.Atoi(id)
	book.ID = uint(Id)
	err = DB.GetInstance().GetDb().Create(&book).Error
	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (ur *BookRepository) SoftDelete(id string) error {
	if err := DB.GetInstance().GetDb().Delete(&model.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (ur *BookRepository) HardDelete(id string) error {
	if err := DB.GetInstance().GetDb().Unscoped().Delete(&model.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}
