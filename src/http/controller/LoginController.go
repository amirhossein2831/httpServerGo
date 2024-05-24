package controller

import (
	"encoding/json"
	"errors"
	"github.com/amirhossein2831/httpServerGo/src/Auth"
	"github.com/amirhossein2831/httpServerGo/src/http/Response"
	"github.com/amirhossein2831/httpServerGo/src/http/repositories"
	"github.com/amirhossein2831/httpServerGo/src/http/request"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials request.CredRequest
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		Response.NewJson().
			SetData(err).
			Log().
			Send(w)
		return
	}
	credentials, err = credentials.Validate()
	if err != nil {
		Response.NewJson().
			SetData(err).
			Log().
			Send(w)
		return
	}

	auth, err := authenticate(credentials.Email, credentials.Password)
	if err != nil {
		Response.NewJson().
			SetData(err).
			Log().
			Send(w)
		return
	}
	if !auth {
		Response.NewJson().
			SetData(errors.New("email or password is wrong")).
			Log().
			Send(w)
		return
	}

	token, err := Auth.CreateToken(credentials.Email)
	if err != nil {
		Response.NewJson().
			SetData(errors.New("email or password is wrong")).
			Log().
			Send(w)
		return
	}

	Response.NewJson().
		SetStatusCode(http.StatusOK).
		SetSuccess(true).
		SetData(map[string]string{"token": token}).
		Log().
		Send(w)
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
