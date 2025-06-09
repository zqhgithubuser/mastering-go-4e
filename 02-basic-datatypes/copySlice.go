package main

import "fmt"

func main() {
	a1 := []int{1}
	a2 := []int{-1, -2}
	a5 := []int{10, 11, 12, 13, 14}
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
	fmt.Println("a5", a5)

	// copy a2 to a1
	copy(a1, a2)
	fmt.Println("a1", a1) // a1 [-1]
	fmt.Println("a2", a2) // a2 [-1 -2]

	// copy a5 to a1
	copy(a1, a5)
	fmt.Println("a1", a1) // a1 [10]
	fmt.Println("a5", a5) // [10 11 12 13 14]

	// copy a2 to a5
	copy(a5, a2)
	fmt.Println("a2", a2)
	fmt.Println("a5", a5) // a5 [-1 -2 12 13 14]
}
