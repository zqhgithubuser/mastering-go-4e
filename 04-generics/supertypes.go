package main

import "fmt"

type AnotherInt int

// AllInts 支持所有底层为 int 的类型
type AllInts interface {
	~int
}

func AddElements[T AllInts](s []T) T {
	sum := T(0)
	for _, v := range s {
		sum += v
	}
	return sum
}

func main() {
	s := []AnotherInt{0, 1, 2}
	fmt.Println(AddElements(s)) // 3
}
