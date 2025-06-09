package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		log.Fatal("Fatal: Hello World!")
	}
	// Panic包含额外的低级别信息
	log.Panic("Panic: Hello World!")
}
