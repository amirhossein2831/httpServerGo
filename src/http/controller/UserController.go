package controller

import (
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/http/service"
	"github.com/amirhossein2831/httpServerGo/src/util"
	"github.com/gorilla/mux"
	"net/http"
)

type UserController struct {
	service *service.UserService
	crud    Crud
}

func NewUserController() *UserController {
	return &UserController{
		service: &service.UserService{},
	}
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := c.service.Get()
	if err != nil {
		util.JsonError(w, err)
		return
	}
	util.JsonResponse(w, http.StatusOK, users)
}

func (c *UserController) Show(w http.ResponseWriter, r *http.Request) {
	user, err := c.service.GetEntity(mux.Vars(r)["id"])
	if err != nil {
		util.JsonError(w, err)
		return
	}

	util.JsonResponse(w, http.StatusOK, user)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest request.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		util.JsonError(w, err)
		return
	}
	user, err := c.service.Post(userRequest)
	if err != nil {
		util.JsonError(w, err)
		return
	}

	util.JsonResponse(w, http.StatusCreated, user)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var userRequest request.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		util.JsonError(w, err)
		return
	}

	user, err := c.service.Put(userRequest, mux.Vars(r)["id"])
	if err != nil {
		util.JsonError(w, err)
		return
	}
	
	util.JsonResponse(w, http.StatusOK, user)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	err := c.service.Delete(mux.Vars(r)["id"])
	if err != nil {
		util.JsonError(w, err)
		return
	}
	util.JsonResponse(w, http.StatusOK, struct{}{})
}
