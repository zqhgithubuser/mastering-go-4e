package main

import (
	"fmt"
	"os"
)

func addFloats(message string, s ...float64) float64 {
	fmt.Println(message)
	sum := float64(0)
	for _, a := range s {
		sum += a
	}
	s[0] = -1000
	return sum
}

func everything(input ...interface{}) {
	fmt.Println(input)
}

func main() {
	sum := addFloats("Adding numbers...", 1.1, 2.12, 3.14, 4, 5, -1, 10)
	fmt.Println("Sum:", sum)

	s := []float64{1.1, 2.12, 3.14}
	sum = addFloats("Adding numbers...", s...)
	fmt.Println("Sum:", sum)
	everything(s)

	empty := make([]interface{}, len(os.Args[1:]))
	for i, v := range os.Args[1:] {
		empty[i] = v
	}
	everything(empty...)
}
