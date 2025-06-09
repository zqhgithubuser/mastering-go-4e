package main

import "fmt"

func main() {
	aString := "Hello World! €"
	fmt.Println("First byte", string(aString[0])) // H

	r := '€'
	fmt.Println("As an int32 value:", r)
	// Convert Runes to text
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)
	for _, v := range aString {
		fmt.Printf("%x ", v) // 48 65 6c 6c 6f 20 57 6f 72 6c 64 21 20 20ac
	}
	fmt.Println()
	for _, v := range aString {
		fmt.Printf("%c ", v) // H e l l o   W o r l d !   €
	}
	fmt.Println()
}
