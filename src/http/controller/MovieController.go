package controller

import (
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/http/service"
	"github.com/amirhossein2831/httpServerGo/src/util"
	"github.com/gorilla/mux"
	"net/http"
)

type MovieController struct {
	service service.Service
}

func NewMovieController() *MovieController {
	return &MovieController{
		service: service.NewMovieService(),
	}
}

func (c *MovieController) Index(w http.ResponseWriter, r *http.Request) {
	movies, err := c.service.Index()
	if err != nil {
		util.JsonError(w, err)
		return
	}

	util.JsonResponse(w, http.StatusOK, movies)
}

func (c *MovieController) Show(w http.ResponseWriter, r *http.Request) {
	movie, err := c.service.Show(mux.Vars(r)["id"])
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

	movie, err := c.service.Create(&movieRequest)
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

	movie, err := c.service.Update(&movieRequest, mux.Vars(r)["id"])
	if err != nil {
		util.JsonError(w, err)
		return
	}

	util.JsonResponse(w, http.StatusOK, movie)
}

func (c *MovieController) Delete(w http.ResponseWriter, r *http.Request) {
	err := c.service.Delete(mux.Vars(r)["id"])
	if err != nil {
		util.JsonError(w, err)
		return
	}
	util.JsonResponse(w, http.StatusOK, struct{}{})
}
