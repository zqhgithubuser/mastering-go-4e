package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("You are using ", runtime.Compiler, " ")    // gc
	fmt.Println("on a", runtime.GOARCH, "machine")        // amd64
	fmt.Println("Using Go version", runtime.Version())    // go1.21.4
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0)) // 8
}
