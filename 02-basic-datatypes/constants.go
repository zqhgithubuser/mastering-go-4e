package main

import (
	"fmt"
)

type Digit int
type Power2 int

const PI = 3.1415926

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1) // 1476
	fmt.Println(PI) // 3.1415926
	// iota: 常量生成器
	const (
		Zero  Digit = iota // 0
		One                // 1
		Two                // 2
		Three              // 3
		Four               // 4
	)
	fmt.Println(One)
	fmt.Println(Two)

	const (
		p2_0 Power2 = 1 << iota // 1 << 0 = 0
		_
		p2_2 // 1 << 2 = 4
		_
		p2_4 // 1 << 4 = 16
		_
		p2_6 // 1 << 6 = 64
	)
	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)
}
