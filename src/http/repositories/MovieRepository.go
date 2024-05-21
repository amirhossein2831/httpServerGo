package repositories

import (
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"strconv"
)

type MovieRepository struct {
}

func NewMovieRepository() *MovieRepository {
	return &MovieRepository{}
}

func (ur *MovieRepository) All() ([]model.Mod, error) {
	var movies []model.Movie

	err := DB.GetInstance().GetDb().Find(&movies).Error
	if err != nil {
		return nil, err
	}

	return model.MovieToMod(movies), nil
}

func (ur *MovieRepository) Get(id string) (model.Mod, error) {
	var movie model.Movie

	err := DB.GetInstance().GetDb().First(&movie, id).Error
	if err != nil {
		return model.Movie{}, err
	}

	return movie, nil
}

func (ur *MovieRepository) GetByColumn(email string) (model.Mod, error) {
	return nil, nil
}

func (ur *MovieRepository) Create(data model.Mod) (model.Mod, error) {
	movie := data.(model.Movie)

	err := DB.GetInstance().GetDb().Create(&movie).Error
	if err != nil {
		return model.Movie{}, err
	}

	return movie, nil
}

func (ur *MovieRepository) Update(data model.Mod, id string) (model.Mod, error) {
	movie := data.(model.User)
	err := ur.HardDelete(id)
	if err != nil {
		return model.Movie{}, err
	}

	Id, _ := strconv.Atoi(id)
	movie.ID = uint(Id)
	err = DB.GetInstance().GetDb().Create(&movie).Error
	if err != nil {
		return model.Movie{}, err
	}

	return movie, nil
}

func (ur *MovieRepository) SoftDelete(id string) error {
	if err := DB.GetInstance().GetDb().Delete(&model.Movie{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (ur *MovieRepository) HardDelete(id string) error {
	if err := DB.GetInstance().GetDb().Unscoped().Delete(&model.Movie{}, id).Error; err != nil {
		return err
	}
	return nil
}
