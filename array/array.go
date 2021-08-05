package main

import "fmt"

func main() {
	var arr1 = [3]int{}
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	arr2 := [3]int{1, 2, 3}

	var arr3 = make([]int, 3)
	arr3[0] = 1
	arr3[1] = 2
	arr3[2] = 3

	// arr1 = append(arr1, 10)

	fmt.Println("arr1", arr1)
	fmt.Println("arr2", arr2)
	fmt.Println("arr3", arr3)

	fmt.Println("____________")

	// ? [start:end]
	slice1 := arr1[:]
	fmt.Println("slice1", slice1)
	fmt.Println(cap(slice1))
	// slice1 = append(slice1, 5)
	// fmt.Println(cap(slice1))

	slice2 := []int{1, 2, 3, 4}
	fmt.Println("slice2", slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	slice2 = append(slice2, 5)
	fmt.Println("____________")
	fmt.Println("slice2", slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	slice2 = append(slice2, 6)
	fmt.Println("____________")
	fmt.Println("slice2", slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
}
