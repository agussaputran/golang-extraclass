package main

import "fmt"

func main() {
	var baru = hello2("agus")
	fmt.Println(baru)

	var name, _ = hello3("string", 23)

	fmt.Println(name)

	listNumber(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// closure function
	var closure1 = func(name string) string {
		// fmt.Println("ini closure ", "hello", name)
		return name
	}

	nama := closure1("agus")
	fmt.Println("nama dari closure", nama)
}

// function with parameter
func hello1(name string) {
	fmt.Println("Hello", name)
}

// function with param and return value
func hello2(name string) string {
	// fmt.Println("Hello", name)
	greetings := "hello " + name
	return greetings
}

// function with params and multiple return value
func hello3(name string, age int) (string, int) {
	return name, age
}

// variadic function
func listNumber(number ...int) {
	fmt.Println(number)
}
