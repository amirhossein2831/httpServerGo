package repositories

import (
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"net/http"
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

func CreateMovie(r *http.Request) (model.Movie, error) {
	var movie model.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		return movie, err
	}
	res := DB.GetInstance().GetDb().Create(&movie)
	if res.Error != nil {
		return movie, res.Error
	}

	return movie, nil
}

func UpdateMovie(r *http.Request, id string) (model.Movie, error) {
	var movie model.Movie
	err := DB.GetInstance().GetDb().First(&movie, id).Error
	if err != nil {
		return movie, err
	}
	err = json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		return movie, err
	}
	DB.GetInstance().GetDb().Save(&movie)

	return movie, nil
}

func DeleteMovie(id string) error {
	if err := DB.GetInstance().GetDb().Delete(&model.Movie{}, id).Error; err != nil {
		return err
	}
	return nil
}
