package Http

import (
	"fmt"
	"github.com/amirhossein2831/httpServerGo/src/Model"
	"log"
	"net/http"
)

var users []Model.User

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

func allUser(writer http.ResponseWriter, r *http.Request) {

}

func getUser(writer http.ResponseWriter, r *http.Request) {

}

func createUser(writer http.ResponseWriter, r *http.Request) {

}

func updateUser(writer http.ResponseWriter, r *http.Request) {

}

func partialUpdateUser(writer http.ResponseWriter, r *http.Request) {

}

func deleteUser(writer http.ResponseWriter, r *http.Request) {

}
