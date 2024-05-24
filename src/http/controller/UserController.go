package controller

import (
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/http/Response"
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/http/service"
	"github.com/gorilla/mux"
	"net/http"
)

type UserController struct {
	service service.Service
}

func NewUserController() *UserController {
	return &UserController{
		service: service.NewUserService(),
	}
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := c.service.Index()
	if err != nil {
		Response.NewJson().
			SetData(err).
			Log().
			Send(w)
		return
	}

	Response.NewJson().
		SetStatusCode(http.StatusOK).
		SetSuccess(true).
		SetData(users).
		Log().
		Send(w)
}

func (c *UserController) Show(w http.ResponseWriter, r *http.Request) {
	user, err := c.service.Show(mux.Vars(r)["id"])
	if err != nil {
		Response.NewJson().
			SetData(err).
			Log().
			Send(w)
		return
	}

	Response.NewJson().
		SetStatusCode(http.StatusOK).
		SetSuccess(true).
		SetData(user).
		Log().
		Send(w)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest request.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		Response.NewJson().
			SetData(err).
			Log().
			Send(w)
		return
	}
	user, err := c.service.Create(&userRequest)
	if err != nil {
		Response.NewJson().
			SetData(err).
			Log().
			Send(w)
		return
	}

	Response.NewJson().
		SetStatusCode(http.StatusCreated).
		SetSuccess(true).
		SetData(user).
		Log().
		Send(w)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var userRequest request.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		Response.NewJson().
			SetData(err).
			Log().
			Send(w)
		return
	}

	user, err := c.service.Update(&userRequest, mux.Vars(r)["id"])
	if err != nil {
		Response.NewJson().
			SetData(err).
			Log().
			Send(w)
		return
	}

	Response.NewJson().
		SetStatusCode(http.StatusOK).
		SetSuccess(true).
		SetData(user).
		Log().
		Send(w)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	err := c.service.Delete(mux.Vars(r)["id"])
	if err != nil {
		Response.NewJson().
			SetData(err).
			Log().
			Send(w)
		return
	}
	Response.NewJson().
		SetStatusCode(http.StatusOK). // i don't know why but panic on 204 (no content)
		SetSuccess(true).
		SetData(struct{}{}).
		Log().
		Send(w)
}
