package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for i := 10; i >= 1; i-- {
		fmt.Println("perulangan ke-", i)
	}

	for _, value := range arr1 {
		if value == 1 {
			fmt.Println("benar")
			continue
		}
	}

	for {
		fmt.Println("perulangan")
		if 1 == 1 {
			break
		}
	}
}
