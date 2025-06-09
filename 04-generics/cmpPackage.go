package main

import (
	"cmp"
	"fmt"
)

func main() {
	// x == y -> 0
	// x > y -> 1
	// x < y -> -1
	fmt.Println(cmp.Compare(5, 4))
	fmt.Println(cmp.Compare(4, 5))
	// x < y -> true
	fmt.Println(cmp.Less(4, 5))
}
