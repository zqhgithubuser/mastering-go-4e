package main

import (
	"fmt"
	"strconv"
)

type record struct {
	Field1 int
	Field2 string
}

func main() {
	var s []record
	for i := 0; i < 10; i++ {
		text := "text" + strconv.Itoa(i)
		temp := record{Field1: i, Field2: text}
		s = append(s, temp)
	}

	fmt.Println("Index 0:", s[0].Field1, s[0].Field2) // Index 0: 0 text0
	fmt.Println("Number of structures:", len(s))      // 10

	sum := 0
	for _, k := range s {
		sum += k.Field1
	}
	fmt.Println("Sum:", sum) // Sum: 45
}
