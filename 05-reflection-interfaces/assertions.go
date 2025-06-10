package main

import "fmt"

func returnNumber() interface{} {
	return 12
}

func main() {
	anInt := returnNumber()
	number, ok := anInt.(int)
	if ok {
		fmt.Println("Type assertion successful: ", number)
	} else {
		fmt.Println("Type assertion failed!")
	}
	number++
	fmt.Println(number) // 13

	value, ok := anInt.(int64)
	if ok {
		fmt.Println("Type assertion successful: ", value)
	} else {
		fmt.Println("Type assertion failed!") // Type assertion failed!
	}

	_ = anInt.(bool) // panic: interface conversion: interface {} is int, not bool
}
