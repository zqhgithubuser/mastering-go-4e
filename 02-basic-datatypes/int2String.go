package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Print provide an integer.")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// int -> string
	input := strconv.Itoa(n)
	fmt.Printf("strconv.Itoa() %s of type %T\n", input, input)

	// int -> string
	input = strconv.FormatInt(int64(n), 10)
	fmt.Printf("strconv.FormatInt() %s of type %T\n", input, input)

	// int -> rune
	input = string(rune(n))
	fmt.Println(input)
	fmt.Printf("string() %s of type %T\n", input, input)
}
