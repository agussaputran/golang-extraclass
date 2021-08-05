package main

import (
	"fmt"
	"strconv"
)

func main() {
	var string1 string = "10"
	var string2 string = "5"

	numStr1, err := strconv.ParseInt(string1, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	numStr2, _ := strconv.ParseInt(string2, 10, 32)

	floatStr1, _ := strconv.ParseFloat(string1, 64)

	fmt.Println(floatStr1)

	result := numStr1 + numStr2
	fmt.Println(result)
}
