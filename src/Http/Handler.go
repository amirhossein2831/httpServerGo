package Http

import (
	"encoding/json"
	"fmt"
	"github.com/amirhossein2831/httpServerGo/src/Model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusBadRequest)
		return
	}
	println("hello url")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
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

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := Model.GetUsers()
	JsonResponse(w, http.StatusOK, users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := Model.GetUser(vars["id"])
	if err != nil {
		JsonError(w, err)
		return
	}
	JsonResponse(w, http.StatusOK, user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user Model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		JsonError(w, err)
		return
	}
	Model.CreateUser(user)
	JsonResponse(w, http.StatusOK, user)
}
