package main

import "fmt"

func main() {
	aMap := map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
	}
	fmt.Println(aMap)

	keys := make([]string, 0, len(aMap))
	values := make([]int, 0, len(aMap))
	for k, v := range aMap {
		keys = append(keys, k)
		values = append(values, v)
	}
	fmt.Println("keys:", keys)
	fmt.Println("values:", values)
}
