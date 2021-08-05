package main

import (
	"fmt"
	"log"
)

func main() {
	// fmt.Println("Hello World")
	// log.Println("Hello world")

	const var1 = "" // const value

	var varNoValue string //? deklarasi variabel tanpa nilai

	var var2 string = "Hello" //? deklarasi variabel 1
	fmt.Println(var2)

	var3 := 3 // ? deklarasi variabel 2
	fmt.Println(var3)

	var2 = "hello lagi"
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(varNoValue)

	log.Println(var1)

}
