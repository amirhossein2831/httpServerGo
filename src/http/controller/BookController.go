package controller

import (
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/repositories"
	"github.com/amirhossein2831/httpServerGo/src/util"
	"github.com/gorilla/mux"
	"net/http"
)

type BookController struct {
	Crud
}

func (c *BookController) Index(w http.ResponseWriter, r *http.Request) {
	books, err := repositories.GetBooks()
	if err != nil {
		util.JsonError(w, err)
		return
	}

	util.JsonResponse(w, http.StatusOK, books)
}

func (c *BookController) Show(w http.ResponseWriter, r *http.Request) {
	book, err := repositories.GetBook(mux.Vars(r)["id"])
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

	book, err := bookRequest.Validate()
	if err != nil {
		util.JsonError(w, err)
		return
	}
	book, err = repositories.CreateBook(book)
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

	book, err := bookRequest.Validate()
	if err != nil {
		util.JsonError(w, err)
		return
	}

	book, err = repositories.UpdateBook(book, mux.Vars(r)["id"])
	if err != nil {
		util.JsonError(w, err)
		return
	}
	util.JsonResponse(w, http.StatusOK, book)
}

func (c *BookController) Delete(w http.ResponseWriter, r *http.Request) {
	err := repositories.SoftDeleteBook(mux.Vars(r)["id"])
	if err != nil {
		util.JsonError(w, err)
		return
	}
	util.JsonResponse(w, http.StatusOK, struct{}{})
}
