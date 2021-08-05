package main

import "fmt"

const number int = 10

func main() {
	fmt.Println("OPERATOR ARITMATIKA")
	var penjumlahan int = 10 + 10
	fmt.Println("penjumlahan", penjumlahan)

	var pengurangan int = 10 - 5
	var perkalian int = 10 * 2

	num1 := 10
	num2 := 3
	var result float32 = float32(num1) / float32(num2)
	result2 := float32(num1 / num2)
	fmt.Println("pembagian2", result2)

	fmt.Println("pengurangan", pengurangan)
	fmt.Println("pembagian", result)
	fmt.Println("perkalian", perkalian)

	num3 := 10
	num3 /= 5

	num4 := 5
	num4--

	fmt.Println(num3)
	fmt.Println(num4)

	var modulo int = 10 % 5
	fmt.Println("hasil bagi", modulo)

	fmt.Println()
	fmt.Println("OPERATOR LOGIKA")

	perbandingan1 := 10 != 10
	var perbandingan2 bool = 10 < 9
	var perbandingan3 bool = 10 <= 9

	fmt.Println(perbandingan1)
	fmt.Println(perbandingan2)
	fmt.Println(perbandingan3)

	operator()
}

func operator() {
	result := number + 10
	fmt.Println(result)
}
