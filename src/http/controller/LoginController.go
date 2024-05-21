package controller

import (
	"encoding/json"
	"errors"
	"github.com/amirhossein2831/httpServerGo/src/Auth"
	"github.com/amirhossein2831/httpServerGo/src/http/repositories"
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"github.com/amirhossein2831/httpServerGo/src/util"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials request.CredRequest
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		util.JsonError(w, err)
		return
	}
	credentials, err = credentials.Validate()
	if err != nil {
		util.JsonError(w, err)
		return
	}

	auth, err := authenticate(credentials.Email, credentials.Password)
	if err != nil {
		util.JsonError(w, err)
		return
	}
	if !auth {
		util.JsonError(w, errors.New("email or password is wrong"))
		return
	}

	token, err := Auth.CreateToken(credentials.Email)
	if err != nil {
		util.JsonError(w, errors.New("email or password is wrong"))
		return
	}

	util.JsonResponse(w, http.StatusOK, map[string]string{"token": token})
}

func authenticate(email string, password string) (bool, error) {
	mod, err := repositories.NewUserRepository().GetByColumn("email", email)
	user := mod.(model.User)
	if err != nil {
		return false, errors.New("can't find the user with given email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, errors.New("the password is not correct")
	}

	return true, nil
}
