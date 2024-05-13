package controller

import (
	http2 "github.com/amirhossein2831/httpServerGo/src/http"
	"github.com/amirhossein2831/httpServerGo/src/repositories"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err, _ := repositories.GetUsers()
	if err != nil {
		http2.JsonResponse(w, http.StatusBadRequest, err)
	}

	http2.JsonResponse(w, http.StatusOK, users)
}

