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
		http2.JsonResponse(w, http.StatusBadRequest, err)
	}

	http2.JsonResponse(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user, err, _ := repositories.GetUser(mux.Vars(r)["id"])
	if err != nil {
		http2.JsonResponse(w, http.StatusBadRequest, err)
	}

	http2.JsonResponse(w, http.StatusOK, user)
}
