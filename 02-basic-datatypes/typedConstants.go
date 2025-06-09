package main

import "fmt"

const (
	typedConstant   = int16(100)
	untypedConstant = 100
)

func main() {
	i := 1
	fmt.Println("unTyped:", i*untypedConstant)    // 100
	fmt.Println("Typed:", int16(i)*typedConstant) // 100
}
