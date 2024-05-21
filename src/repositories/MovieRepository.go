package repositories

import (
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"strconv"
)

func GetMovies() ([]model.Movie, error) {
	var movies []model.Movie
	res := DB.GetInstance().GetDb().Find(&movies)
	if res.Error != nil {
		return nil, res.Error
	}
	return movies, nil
}

func GetMovie(id string) (model.Movie, error) {
	var movie model.Movie

	res := DB.GetInstance().GetDb().First(&movie, id)
	if res.Error != nil {
		return movie, res.Error
	}
	return movie, nil
}

func CreateMovie(movie model.Movie) (model.Movie, error) {
	res := DB.GetInstance().GetDb().Create(&movie)
	if res.Error != nil {
		return model.Movie{}, res.Error
	}

	return movie, nil
}

func UpdateMovie(movie model.Movie, id string) (model.Movie, error) {
	err := HardDeleteMovie(id)
	if err != nil {
		return model.Movie{}, err
	}

	Id, _ := strconv.Atoi(id)
	movie.ID = uint(Id)
	res := DB.GetInstance().GetDb().Create(&movie)
	if res.Error != nil {
		return model.Movie{}, res.Error
	}

	return movie, nil
}

func SoftDeleteMovie(id string) error {
	if err := DB.GetInstance().GetDb().Delete(&model.Movie{}, id).Error; err != nil {
		return err
	}
	return nil
}

func HardDeleteMovie(id string) error {
	if err := DB.GetInstance().GetDb().Unscoped().Delete(&model.Movie{}, id).Error; err != nil {
		return err
	}
	return nil
}
