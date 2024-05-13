package http

import (
	"encoding/json"
	"fmt"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusBadRequest)
		return
	}
	println("hello url")
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusBadRequest)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Printf("there is a error while parsing the form %v", err)
		return
	}
	fmt.Printf("the name is %v \n", r.FormValue("name"))
	fmt.Printf("the address is %v \n", r.FormValue("address"))

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := model.GetUsers()
	JsonResponse(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := model.GetUser(vars["id"])
	if err != nil {
		JsonError(w, err)
		return
	}
	JsonResponse(w, http.StatusOK, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		JsonError(w, err)
		return
	}
	model.CreateUser(user)
	JsonResponse(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var body model.User
	vars := mux.Vars(r)
	user, err := model.GetUser(vars["id"])
	if err != nil {
		JsonError(w, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		JsonError(w, err)
		return
	}

	model.UpdateUser(user, body)
	JsonResponse(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	err := model.RemoveUser(vars["id"])
	if err != nil {
		JsonError(w, err)
		return
	}
	JsonResponse(w, http.StatusOK, struct{}{})
}
