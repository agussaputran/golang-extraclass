package main

import "fmt"

func main() {
	map1 := map[string]string{
		"name": "agus",
		"age":  "24",
	}

	map2 := map[string]string{
		"name": "agus",
		"age":  "24",
	}

	map3 := map[string][]interface{}{
		"name": {"agus", "saputra"},
		"age":  {24},
	}

	arr1 := []map[string]string{map1, map2}
	fmt.Println(arr1)

	arr2 := []interface{}{}

	arr2 = append(arr2, map1, map2, map3)
	fmt.Println("ini arr2", arr2)

	fmt.Println(map1["age"])
	fmt.Println(map3)
}
