package controller

import (
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/http/Response"
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/http/service"
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
		Response.NewJson().
			SetData(err).
			Send(w)
		return
	}

	Response.NewJson().
		SetStatusCode(http.StatusOK).
		SetSuccess(true).
		SetData(movies).
		Send(w)
}

func (c *MovieController) Show(w http.ResponseWriter, r *http.Request) {
	movie, err := c.service.Show(mux.Vars(r)["id"])
	if err != nil {
		Response.NewJson().
			SetData(err).
			Send(w)
		return
	}

	Response.NewJson().
		SetStatusCode(http.StatusOK).
		SetSuccess(true).
		SetData(movie).
		Send(w)
}

func (c *MovieController) Create(w http.ResponseWriter, r *http.Request) {
	var movieRequest request.MovieRequest
	err := json.NewDecoder(r.Body).Decode(&movieRequest)
	if err != nil {
		Response.NewJson().
			SetData(err).
			Send(w)
		return
	}

	movie, err := c.service.Create(&movieRequest)
	if err != nil {
		Response.NewJson().
			SetData(err).
			Send(w)
		return
	}

	Response.NewJson().
		SetStatusCode(http.StatusOK).
		SetSuccess(true).
		SetData(movie).
		Send(w)
}

func (c *MovieController) Update(w http.ResponseWriter, r *http.Request) {
	var movieRequest request.MovieRequest
	err := json.NewDecoder(r.Body).Decode(&movieRequest)
	if err != nil {
		Response.NewJson().
			SetData(err).
			Send(w)
		return
	}

	movie, err := c.service.Update(&movieRequest, mux.Vars(r)["id"])
	if err != nil {
		Response.NewJson().
			SetData(err).
			Send(w)
		return
	}

	Response.NewJson().
		SetStatusCode(http.StatusOK).
		SetSuccess(true).
		SetData(movie).
		Send(w)
}

func (c *MovieController) Delete(w http.ResponseWriter, r *http.Request) {
	err := c.service.Delete(mux.Vars(r)["id"])
	if err != nil {
		Response.NewJson().
			SetData(err).
			Send(w)
		return
	}
	Response.NewJson().
		SetStatusCode(http.StatusOK).
		SetSuccess(true).
		SetData(struct{}{}).
		Send(w)
}
