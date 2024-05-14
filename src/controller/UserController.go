package controller

import (
	http2 "github.com/amirhossein2831/httpServerGo/src/http"
	"github.com/amirhossein2831/httpServerGo/src/repositories"
	"github.com/gorilla/mux"
	"net/http"
)

type UserController struct {
	Crud
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := repositories.GetUsers()
	if err != nil {
		http2.JsonError(w, err)
		return
	}

	http2.JsonResponse(w, http.StatusOK, users)
}

func (c *UserController) Show(w http.ResponseWriter, r *http.Request) {
	user, err := repositories.GetUser(mux.Vars(r)["id"])
	if err != nil {
		http2.JsonError(w, err)
		return
	}

	http2.JsonResponse(w, http.StatusOK, user)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	user, err := repositories.CreateUser(r)
	if err != nil {
		http2.JsonError(w, err)
		return
	}
	http2.JsonResponse(w, http.StatusCreated, user)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	user, err := repositories.UpdateUser(r, mux.Vars(r)["id"])
	if err != nil {
		http2.JsonError(w, err)
		return
	}
	http2.JsonResponse(w, http.StatusOK, user)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	err := repositories.DeleteUser(mux.Vars(r)["id"])
	if err != nil {
		http2.JsonError(w, err)
		return
	}
	http2.JsonResponse(w, http.StatusOK, struct{}{})
}
