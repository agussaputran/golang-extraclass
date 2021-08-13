package main

import (
	"belajar-be-dasar/controller"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, This is my first api in go")
	})

	http.HandleFunc("/hello", controller.Hello)
	http.HandleFunc("/get-user", controller.GetUser)
	http.HandleFunc("/post-user", controller.PostUser)

	fmt.Println("server up and run on localhost:9000")
	http.ListenAndServe(":9000", nil)
}
