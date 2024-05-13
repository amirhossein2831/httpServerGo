package controller

import (
	http2 "github.com/amirhossein2831/httpServerGo/src/http"
	"github.com/amirhossein2831/httpServerGo/src/repositories"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err, _ := repositories.GetUsers()
	if err != nil {
		http2.JsonError(w, err)
		return
	}

	http2.JsonResponse(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user, err, _ := repositories.GetUser(mux.Vars(r)["id"])
	if err != nil {
		http2.JsonError(w, err)
		return
	}

	http2.JsonResponse(w, http.StatusOK, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err, _ := repositories.CreateUser(r)
	if err != nil {
		http2.JsonError(w, err)
		return
	}
	http2.JsonResponse(w, http.StatusCreated, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	err := repositories.DeleteUser(mux.Vars(r)["id"])
	if err != nil {
		http2.JsonError(w, err)
		return
	}
	http2.JsonResponse(w, http.StatusOK, struct{}{})
}
