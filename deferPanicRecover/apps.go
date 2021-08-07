package main

import (
	"fmt"
)

// func catch() {
// 	var r = recover()
// 	if r != nil {
// 		fmt.Println("error ", r)
// 	} else {
// 		fmt.Println("Aplikasi berjalan dengan lancar")
// 	}
// }

func main() {
	// defer
	// defer fmt.Println("ini dijalankan terakhir")
	// fmt.Println("ini dijalankan pertama")

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

func insertData() {
	defer func() {
		var r = recover()
		if r != nil {
			fmt.Println("error ", r)
		} else {
			fmt.Println("Aplikasi berjalan dengan lancar")
		}
	}()
}
