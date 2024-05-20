package repositories

import (
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/http/request"
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
	var movieRequest request.MovieRequest
	err := json.NewDecoder(r.Body).Decode(&movieRequest)
	if err != nil {
		return model.Movie{}, err
	}

	movie := movieRequest.ToMovie()
	res := DB.GetInstance().GetDb().Create(&movie)
	if res.Error != nil {
		return model.Movie{}, res.Error
	}

	return movie, nil
}

func UpdateMovie(r *http.Request, id string) (model.Movie, error) {
	var movieRequest request.MovieRequest

	err := DeleteMovie(id)
	if err != nil {
		return model.Movie{}, err
	}

	err = json.NewDecoder(r.Body).Decode(&movieRequest)
	if err != nil {
		return model.Movie{}, err
	}
	movie := movieRequest.ToMovie()
	res := DB.GetInstance().GetDb().Create(&movie)
	if res.Error != nil {
		return model.Movie{}, res.Error
	}

	return movie, nil
}

func DeleteMovie(id string) error {
	if err := DB.GetInstance().GetDb().Delete(&model.Movie{}, id).Error; err != nil {
		return err
	}
	return nil
}
