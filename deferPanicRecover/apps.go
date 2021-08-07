package main

import (
	"fmt"
)

// function untuk recover
func catch() {
	var r = recover()
	if r != nil {
		fmt.Println("error ", r)
	} else {
		fmt.Println("Aplikasi berjalan dengan lancar")
	}
}

func main() {
	// defer
	// defer fmt.Println("ini dijalankan terakhir")
	// fmt.Println("ini dijalankan pertama")

	// cara 1 untuk menjalankan recover dengan memanggil function yang berisi perintah recover
	// defer catch()

	// cara 2 untuk memanggil recover langsung pada function saat ini
	// menggunakan cara IIFE func
	defer func() {
		var r = recover()
		if r != nil {
			fmt.Println("error ", r)
		} else {
			fmt.Println("Aplikasi berjalan dengan lancar")
		}
	}()

	// panic
	panic("error panic")

	// recover
}
