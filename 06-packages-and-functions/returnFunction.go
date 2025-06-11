package main

import "fmt"

func funRet(i int) func(int) int {
	if i < 0 {
		return func(k int) int {
			k = -k
			return k + k
		}
	}
	return func(k int) int {
		return k * k
	}
}

func main() {
	n := 10
	i := funRet(n)
	j := funRet(-4)
	fmt.Printf("%T\n", i) // func(int) int
	fmt.Printf("%T\n", j) // func(int) int
	fmt.Println(j(-5))    // 10
	fmt.Println(i(10))    // 100
	fmt.Println(j(10))    // -20
}
