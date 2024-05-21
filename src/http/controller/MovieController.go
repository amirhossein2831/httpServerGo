package controller

import (
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/repositories"
	"github.com/amirhossein2831/httpServerGo/src/util"
	"github.com/gorilla/mux"
	"net/http"
)

type MovieController struct {
	Crud
}

func (c *MovieController) Index(w http.ResponseWriter, r *http.Request) {
	movies, err := repositories.GetMovies()
	if err != nil {
		util.JsonError(w, err)
		return
	}

	util.JsonResponse(w, http.StatusOK, movies)
}

func (c *MovieController) Show(w http.ResponseWriter, r *http.Request) {
	movie, err := repositories.GetMovie(mux.Vars(r)["id"])
	if err != nil {
		util.JsonError(w, err)
		return
	}

	util.JsonResponse(w, http.StatusOK, movie)
}

func (c *MovieController) Create(w http.ResponseWriter, r *http.Request) {
	var movieRequest request.MovieRequest
	err := json.NewDecoder(r.Body).Decode(&movieRequest)
	if err != nil {
		util.JsonError(w, err)
		return
	}

	movie, err := movieRequest.Validate()
	if err != nil {
		util.JsonError(w, err)
		return
	}

	movie, err = repositories.CreateMovie(movie)
	if err != nil {
		util.JsonError(w, err)
		return
	}
	util.JsonResponse(w, http.StatusCreated, movie)
}

func (c *MovieController) Update(w http.ResponseWriter, r *http.Request) {
	var movieRequest request.MovieRequest
	err := json.NewDecoder(r.Body).Decode(&movieRequest)
	if err != nil {
		util.JsonError(w, err)
		return
	}

	movie, err := movieRequest.Validate()
	if err != nil {
		util.JsonError(w, err)
		return
	}

	movie, err = repositories.UpdateMovie(movie, mux.Vars(r)["id"])
	if err != nil {
		util.JsonError(w, err)
		return
	}
	util.JsonResponse(w, http.StatusOK, movie)
}

func (c *MovieController) Delete(w http.ResponseWriter, r *http.Request) {
	err := repositories.SoftDeleteMovie(mux.Vars(r)["id"])
	if err != nil {
		util.JsonError(w, err)
		return
	}
	util.JsonResponse(w, http.StatusOK, struct{}{})
}
