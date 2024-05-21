package controller

import (
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/http/service"
	"github.com/amirhossein2831/httpServerGo/src/util"
	"github.com/gorilla/mux"
	"net/http"
)

type BookController struct {
	service service.Service
	Crud
}

func NewBookController() *BookController {
	return &BookController{
		service: service.NewMovieService(),
	}
}

func (c *BookController) Index(w http.ResponseWriter, r *http.Request) {
	books, err := c.service.Index()
	if err != nil {
		util.JsonError(w, err)
		return
	}

	util.JsonResponse(w, http.StatusOK, books)
}

func (c *BookController) Show(w http.ResponseWriter, r *http.Request) {
	book, err := c.service.Show(mux.Vars(r)["id"])
	if err != nil {
		util.JsonError(w, err)
		return
	}

	util.JsonResponse(w, http.StatusOK, book)
}

func (c *BookController) Create(w http.ResponseWriter, r *http.Request) {
	var bookRequest request.BookRequest
	err := json.NewDecoder(r.Body).Decode(&bookRequest)
	if err != nil {
		util.JsonError(w, err)
		return
	}

	book, err := c.service.Create(&bookRequest)
	if err != nil {
		util.JsonError(w, err)
		return
	}

	util.JsonResponse(w, http.StatusCreated, book)
}

func (c *BookController) Update(w http.ResponseWriter, r *http.Request) {
	var bookRequest request.BookRequest
	err := json.NewDecoder(r.Body).Decode(&bookRequest)
	if err != nil {
		util.JsonError(w, err)
		return
	}

	book, err := c.service.Update(&bookRequest, mux.Vars(r)["id"])
	if err != nil {
		util.JsonError(w, err)
		return
	}

	util.JsonResponse(w, http.StatusOK, book)
}

func (c *BookController) Delete(w http.ResponseWriter, r *http.Request) {
	err := c.service.Delete(mux.Vars(r)["id"])
	if err != nil {
		util.JsonError(w, err)
		return
	}
	util.JsonResponse(w, http.StatusOK, struct{}{})
}
