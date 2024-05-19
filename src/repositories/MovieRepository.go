package repositories

import (
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/App"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"net/http"
)

func GetMovies() ([]model.Movie, error) {
	var movies []model.Movie
	res := App.App.GetDB().Find(&movies)
	if res.Error != nil {
		return nil, res.Error
	}
	return movies, nil
}

func GetMovie(id string) (model.Movie, error) {
	var movie model.Movie

	res := App.App.GetDB().First(&movie, id)
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
	res := App.App.GetDB().Create(&movie)
	if res.Error != nil {
		return movie, res.Error
	}

	return movie, nil
}

func UpdateMovie(r *http.Request, id string) (model.Movie, error) {
	var movie model.Movie
	err := App.App.GetDB().First(&movie, id).Error
	if err != nil {
		return movie, err
	}
	err = json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		return movie, err
	}
	App.App.GetDB().Save(&movie)

	return movie, nil
}

func DeleteMovie(id string) error {
	if err := App.App.GetDB().Delete(&model.Movie{}, id).Error; err != nil {
		return err
	}
	return nil
}
