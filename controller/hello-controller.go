package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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

func SingleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	// err := r.ParseMultipartForm(1024)
	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		fmt.Println("Error ParseMultipartForm", err)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		fmt.Println("Error on get FormFile", err)
		return
	}

	// fileType := handler.Header.Get("Content-Type")

	basePath, err := os.Getwd()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		fmt.Println("error on get path")
		return
	}

	fileLocation := filepath.Join(basePath, "files", handler.Filename)

	dst, err := os.Create(fileLocation)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		fmt.Println("Error on create local file", err)
		return
	}

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		fmt.Println("Error on get Copy file", err)
		return
	}

	w.Write([]byte("Upload success"))

}

func MultipleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	basePath, err := os.Getwd()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		fmt.Println("error on get path")
		return
	}

	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		fmt.Println("error on MultipartReader", err)
		return
	}

	for {
		part, err := reader.NextPart()

		if err == io.EOF {
			break
		}

		fileLocation := filepath.Join(basePath, "files", part.FileName())
		dst, err := os.Create(fileLocation)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println("Error on create local file", err)
			return
		}

		_, err = io.Copy(dst, part)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			fmt.Println("Error on get Copy file", err)
			return
		}
	}
	w.Write([]byte("Upload success"))

}
