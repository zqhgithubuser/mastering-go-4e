package main

import "fmt"

func main() {
	b := make([]byte, 12)
	fmt.Println("Byte slice:", b)
	b = []byte("Byte slice €")
	fmt.Println("Byte slice:", b) // [66 121 116 101 32 115 108 105 99 101 32 226 130 172]

	fmt.Printf("Byte slice as text: %s\n", b)     // Byte slice €
	fmt.Println("Byte slice as text:", string(b)) // Byte slice €

	fmt.Println("Length of b:", len(b))
}
