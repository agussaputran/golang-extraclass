package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var v1 int = 20

	var v2 int = v1

	var v3 *int = &v1 // &v1 itu tidak sama dengan nilai v1 = 20
	// &v1 adalah alamat memori variabel v1

	v1 = 30
	// v3 = 20

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(*v3)

	var p1 = person{"agus", 24}
	var p2 = p1

	var p3 *person = &p1
	p1.name = "fandy"

	p3.name = "haikal"

	fmt.Println("p1", p1)
	fmt.Println("p2", p2)
	fmt.Println("p3", *p3)
}
