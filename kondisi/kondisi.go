package main

import "fmt"

func main() {
	var1 := 1
	var2 := 2

	if var1 != var2 {
		fmt.Println("benar")
	} else {
		fmt.Println("salah")
	}

	switch var1 {
	case 1:
		fmt.Println("benar")
	case 2:
		fmt.Println("salah")
	}

	switch {
	case (var1 == var2):
		fmt.Println("salah")
	case (var2 != var1):
		fmt.Println("benar")
	}

}
