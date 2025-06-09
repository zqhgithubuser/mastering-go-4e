package main

import "fmt"

func main() {
	aMap := map[string]int{}
	fmt.Println("aMap:", aMap) // aMap: map[]

	aMap["test"] = 1
	fmt.Println("aMap:", aMap) // aMap: map[test:1]

	aMap = nil
	fmt.Println("aMap:", aMap) // aMap: map[]
	if aMap == nil {
		fmt.Println("nil map!") // nil map!
		aMap = map[string]int{}
	}

	aMap["test"] = 1

	aMap = nil       // nil map 不用用于写操作
	aMap["test"] = 1 // panic: assignment to entry in nil map
}
