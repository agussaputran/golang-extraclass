package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Username string
	Password string
}

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Bad Request", 400)
		return
	}
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Fprintln(w, "Hello, My Name is", name, "and i'm", age, "years old")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user User
	user.Username = "agussaputran"
	user.Password = "agus34"
	if r.Method != "GET" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	dataByte, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Cannot marshal data", err)
	}

	w.Write(dataByte)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user User

	if r.Method != "POST" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	dataByte, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataByte, &user)
	if err != nil {
		fmt.Println("Cannot Unmarshall data", err)
	}

	fmt.Println(user.Username, user.Password)
	w.Write(dataByte)
}
