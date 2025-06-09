package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Please give me your name: ")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		panic(err)
	}
	fmt.Println("Your name is", name)
}
