package main

import (
	"fmt"
)

func main() {
	c1 := 12 + 1i
	c2 := complex(5, 7)
	fmt.Printf("Type of c1: %T\n", c1) // Type of c1: complex128
	fmt.Printf("Type of c2: %T\n", c2)

	var c3 = complex64(c1 + c2)
	fmt.Println("c3:", c3)
	fmt.Printf("Type of c3: %T\n", c3) // Type of c3: complex64

	x := 12
	k := 5
	fmt.Println(x)
	fmt.Printf("Type of x: %T\n", x) // int
	div := x / k
	fmt.Println("div", div) // 2

	var m, n float64
	m = 1.223
	fmt.Println("m, n:", m, n)

	y := 4 / 2.3
	fmt.Println("y:", y) // 1.7391304347826086

	divFloat := float64(x) / float64(k)
	fmt.Println("divFloat", divFloat)              // 2.4
	fmt.Printf("Type of divFloat: %T\n", divFloat) // float64
}
