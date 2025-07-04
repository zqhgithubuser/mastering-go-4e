package main

import "fmt"

type S1 struct {
	F1 int
	F2 string
}

type S2 struct {
	F1 int
	F2 S1
}

func Print(s interface{}) {
	fmt.Println(s)
}

func main() {
	v1 := S1{10, "Hello"}
	v2 := S2{F1: -1, F2: v1}
	Print(v1) // {10 Hello}
	Print(v2) // {-1 {10 Hello}}

	Print(123)
	Print("Go is the best!")
}
